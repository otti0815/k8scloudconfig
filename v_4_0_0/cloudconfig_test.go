package v_4_0_0

import (
	"encoding/base64"
	"path"
	"testing"

	ignition "github.com/giantswarm/k8scloudconfig/ignition/v_2_2_0"
)

func TestCloudConfig(t *testing.T) {
	tests := []struct {
		name             string
		template         string
		params           Params
		expectedEtcdPort int
	}{
		{
			name:     "master",
			template: MasterTemplate,
			params: Params{
				Extension: nopExtension{},
			},
			expectedEtcdPort: 443,
		},
		{
			name:     "worker",
			template: WorkerTemplate,
			params: Params{
				Extension: nopExtension{},
			},
			expectedEtcdPort: 443,
		},
		{
			name:     "worker",
			template: WorkerTemplate,
			params: Params{
				EtcdPort:  2379,
				Extension: nopExtension{},
			},
			expectedEtcdPort: 2379,
		},
	}

	for _, tc := range tests {
		c := DefaultCloudConfigConfig()

		packagePath, err := GetPackagePath()
		if err != nil {
			t.Errorf("failed to retrieve package path, %v:", err)
		}
		filesPath := path.Join(packagePath, version, filesDir)
		files, err := RenderFiles(filesPath, tc.params)
		if err != nil {
			t.Errorf("failed to render ignition files, %v:", err)
		}
		tc.params.Files = files

		c.Params = tc.params
		c.Template = tc.template

		cloudConfig, err := NewCloudConfig(c)
		if err != nil {
			t.Fatal(err)
		}

		if cloudConfig.params.EtcdPort != tc.expectedEtcdPort {
			t.Errorf("expected etcd port %q, got %q", tc.expectedEtcdPort, cloudConfig.params.EtcdPort)
		}
		if err := cloudConfig.ExecuteTemplate(); err != nil {
			t.Fatal(err)
		}

		cloudConfigBase64 := cloudConfig.Base64()
		if _, err := base64.StdEncoding.DecodeString(cloudConfigBase64); err != nil {
			t.Errorf("The string isn't Base64 encoded: %v", err)
		}

		_, err = ignition.ConvertTemplatetoJSON([]byte(cloudConfig.String()))
		if err != nil {
			t.Fatalf("failed to validate ignition %#q config, %v:", tc.name, err)
		}

	}
}
