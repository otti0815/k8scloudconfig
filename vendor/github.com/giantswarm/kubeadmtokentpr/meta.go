package kubeadmtokentpr

// Metadata to create operatorkit TPR.

const (
	// Name represents the name of the third party resource within Kubernetes.
	Name = "kubeadm-token.giantswarm.io"

	// VersionV1 is the v1 version of this resource.
	VersionV1 = "v1"

	// Description is the description of this resource.
	Description = "The token-operator generates kubeadm tokens."
)
