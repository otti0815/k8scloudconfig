package ignition

// This schema structure is based on github.com/coreos/ignition/config/v2_2/types/schema.go
// Due to issue with unmarshalling embedded anonymous nested structures,
// this file removes such structures.
// Changed types: Directory, File, Link.
type CaReference struct {
	Source       string       `json:"source,omitempty"`
	Verification Verification `json:"verification,omitempty"`
}
type Config struct {
	Ignition Ignition `json:"ignition"`
	Networkd Networkd `json:"networkd,omitempty"`
	Passwd   Passwd   `json:"passwd,omitempty"`
	Storage  Storage  `json:"storage,omitempty"`
	Systemd  Systemd  `json:"systemd,omitempty"`
}
type ConfigReference struct {
	Source       string       `json:"source,omitempty"`
	Verification Verification `json:"verification,omitempty"`
}
type Create struct {
	Force   bool           `json:"force,omitempty"`
	Options []CreateOption `json:"options,omitempty"`
}
type CreateOption string
type Device string
type Directory struct {
	Filesystem string     `json:"filesystem,omitempty"`
	Group      *NodeGroup `json:"group,omitempty"`
	Mode       *int       `json:"mode,omitempty"`
	Overwrite  *bool      `json:"overwrite,omitempty"`
	Path       string     `json:"path,omitempty"`
	User       *NodeUser  `json:"user,omitempty"`
}
type Disk struct {
	Device     string      `json:"device,omitempty"`
	Partitions []Partition `json:"partitions,omitempty"`
	WipeTable  bool        `json:"wipeTable,omitempty"`
}
type File struct {
	Append     bool         `json:"append,omitempty"`
	Contents   FileContents `json:"contents,omitempty"`
	Filesystem string       `json:"filesystem,omitempty"`
	Mode       int          `json:"mode,omitempty"`
	Group      *NodeGroup   `json:"group,omitempty"`
	Overwrite  *bool        `json:"overwrite,omitempty"`
	Path       string       `json:"path,omitempty"`
	User       *NodeUser    `json:"user,omitempty"`
}
type FileContents struct {
	Compression  string       `json:"compression,omitempty"`
	Source       string       `json:"source,omitempty"`
	Verification Verification `json:"verification,omitempty"`
}
type Filesystem struct {
	Mount *Mount  `json:"mount,omitempty"`
	Name  string  `json:"name,omitempty"`
	Path  *string `json:"path,omitempty"`
}
type Group string
type Ignition struct {
	Config   IgnitionConfig `json:"config,omitempty"`
	Security Security       `json:"security,omitempty"`
	Timeouts Timeouts       `json:"timeouts,omitempty"`
	Version  string         `json:"version,omitempty"`
}
type IgnitionConfig struct {
	Append  []ConfigReference `json:"append,omitempty"`
	Replace *ConfigReference  `json:"replace,omitempty"`
}
type Link struct {
	Filesystem string     `json:"filesystem,omitempty"`
	Group      *NodeGroup `json:"group,omitempty"`
	Hard       bool       `json:"hard,omitempty"`
	Overwrite  *bool      `json:"overwrite,omitempty"`
	Path       string     `json:"path,omitempty"`
	Target     string     `json:"target,omitempty"`
	User       *NodeUser  `json:"user,omitempty"`
}
type Mount struct {
	Create         *Create       `json:"create,omitempty"`
	Device         string        `json:"device,omitempty"`
	Format         string        `json:"format,omitempty"`
	Label          *string       `json:"label,omitempty"`
	Options        []MountOption `json:"options,omitempty"`
	UUID           *string       `json:"uuid,omitempty"`
	WipeFilesystem bool          `json:"wipeFilesystem,omitempty"`
}
type MountOption string
type Networkd struct {
	Units []Networkdunit `json:"units,omitempty"`
}
type NetworkdDropin struct {
	Contents string `json:"contents,omitempty"`
	Name     string `json:"name,omitempty"`
}
type Networkdunit struct {
	Contents string           `json:"contents,omitempty"`
	Dropins  []NetworkdDropin `json:"dropins,omitempty"`
	Name     string           `json:"name,omitempty"`
}
type Node struct {
	Filesystem string     `json:"filesystem,omitempty"`
	Group      *NodeGroup `json:"group,omitempty"`
	Overwrite  *bool      `json:"overwrite,omitempty"`
	Path       string     `json:"path,omitempty"`
	User       *NodeUser  `json:"user,omitempty"`
}
type NodeGroup struct {
	ID   *int   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type NodeUser struct {
	ID   *int   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type Partition struct {
	GUID     string `json:"guid,omitempty"`
	Label    string `json:"label,omitempty"`
	Number   int    `json:"number,omitempty"`
	Size     int    `json:"size,omitempty"`
	Start    int    `json:"start,omitempty"`
	TypeGUID string `json:"typeGuid,omitempty"`
}
type Passwd struct {
	Groups []PasswdGroup `json:"groups,omitempty"`
	Users  []PasswdUser  `json:"users,omitempty"`
}
type PasswdGroup struct {
	Gid          *int   `json:"gid,omitempty"`
	Name         string `json:"name,omitempty"`
	PasswordHash string `json:"passwordHash,omitempty"`
	System       bool   `json:"system,omitempty"`
}
type PasswdUser struct {
	Create            *Usercreate        `json:"create,omitempty"`
	Gecos             string             `json:"gecos,omitempty"`
	Groups            []Group            `json:"groups,omitempty"`
	HomeDir           string             `json:"homeDir,omitempty"`
	Name              string             `json:"name,omitempty"`
	NoCreateHome      bool               `json:"noCreateHome,omitempty"`
	NoLogInit         bool               `json:"noLogInit,omitempty"`
	NoUserGroup       bool               `json:"noUserGroup,omitempty"`
	PasswordHash      *string            `json:"passwordHash,omitempty"`
	PrimaryGroup      string             `json:"primaryGroup,omitempty"`
	SSHAuthorizedKeys []SSHAuthorizedKey `json:"sshAuthorizedKeys,omitempty"`
	Shell             string             `json:"shell,omitempty"`
	System            bool               `json:"system,omitempty"`
	UID               *int               `json:"uid,omitempty"`
}
type Raid struct {
	Devices []Device     `json:"devices,omitempty"`
	Level   string       `json:"level,omitempty"`
	Name    string       `json:"name,omitempty"`
	Options []RaidOption `json:"options,omitempty"`
	Spares  int          `json:"spares,omitempty"`
}
type RaidOption string
type SSHAuthorizedKey string
type Security struct {
	TLS TLS `json:"tls,omitempty"`
}
type Storage struct {
	Directories []Directory  `json:"directories,omitempty"`
	Disks       []Disk       `json:"disks,omitempty"`
	Files       []File       `json:"files,omitempty"`
	Filesystems []Filesystem `json:"filesystems,omitempty"`
	Links       []Link       `json:"links,omitempty"`
	Raid        []Raid       `json:"raid,omitempty"`
}
type Systemd struct {
	Units []Unit `json:"units,omitempty"`
}
type SystemdDropin struct {
	Contents string `json:"contents,omitempty"`
	Name     string `json:"name,omitempty"`
}
type TLS struct {
	CertificateAuthorities []CaReference `json:"certificateAuthorities,omitempty"`
}
type Timeouts struct {
	HTTPResponseHeaders *int `json:"httpResponseHeaders,omitempty"`
	HTTPTotal           *int `json:"httpTotal,omitempty"`
}
type Unit struct {
	Contents string          `json:"contents,omitempty"`
	Dropins  []SystemdDropin `json:"dropins,omitempty"`
	Enable   bool            `json:"enable,omitempty"`
	Enabled  *bool           `json:"enabled,omitempty"`
	Mask     bool            `json:"mask,omitempty"`
	Name     string          `json:"name,omitempty"`
}
type Usercreate struct {
	Gecos        string            `json:"gecos,omitempty"`
	Groups       []UsercreateGroup `json:"groups,omitempty"`
	HomeDir      string            `json:"homeDir,omitempty"`
	NoCreateHome bool              `json:"noCreateHome,omitempty"`
	NoLogInit    bool              `json:"noLogInit,omitempty"`
	NoUserGroup  bool              `json:"noUserGroup,omitempty"`
	PrimaryGroup string            `json:"primaryGroup,omitempty"`
	Shell        string            `json:"shell,omitempty"`
	System       bool              `json:"system,omitempty"`
	UID          *int              `json:"uid,omitempty"`
}
type UsercreateGroup string
type Verification struct {
	Hash *string `json:"hash,omitempty"`
}
