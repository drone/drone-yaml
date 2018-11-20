package yaml

// Pipeline is a resource that defines a continuous
// delivery pipeline.
type Pipeline struct {
	Version string `json:"version,omitempty"`
	Kind    string `json:"kind,omitempty"`
	Type    string `json:"type,omitempty"`
	Name    string `json:"name,omitempty"`

	Clone       Clone             `json:"clone,omitempty"`
	Concurrency Concurrency       `json:"concurrency,omitempty"`
	DependsOn   []string          `json:"depends_on,omitempty" yaml:"depends_on" `
	Node        map[string]string `json:"node,omitempty" yaml:"node"`
	Platform    Platform          `json:"platform,omitempty"`
	Services    []*Container      `json:"services,omitempty"`
	Steps       []*Container      `json:"steps,omitempty"`
	Trigger     Conditions        `json:"trigger,omitempty"`
	Volumes     []*Volume         `json:"volumes,omitempty"`
	Workspace   Workspace         `json:"workspace,omitempty"`
}

// GetVersion returns the resource version.
func (p *Pipeline) GetVersion() string { return p.Version }

// GetKind returns the resource kind.
func (p *Pipeline) GetKind() string { return p.Kind }

type (
	// Clone configures the git clone.
	Clone struct {
		Disable bool `json:"disable,omitempty"`
		Depth   int  `json:"depth,omitempty"`
	}

	// Concurrency limits pipeline concurrency.
	Concurrency struct {
		Limit int `json:"limit,omitempty"`
	}

	// Container defines a Docker container configuration.
	Container struct {
		Build       *Build                `json:"build,omitempty"`
		Command     []string              `json:"command,omitempty"`
		Commands    []string              `json:"commands,omitempty"`
		Detach      bool                  `json:"detach,omitempty"`
		DependsOn   []string              `json:"depends_on,omitempty" yaml:"depends_on"`
		Devices     []*VolumeDevice       `json:"devices,omitempty"`
		DNS         []string              `json:"dns,omitempty"`
		DNSSearch   []string              `json:"dns_search,omitempty" yaml:"dns_search"`
		Entrypoint  []string              `json:"entrypoint,omitempty"`
		Environment map[string]*Parameter  `json:"environment,omitempty"`
		ExtraHosts  []string              `json:"extra_hosts,omitempty" yaml:"extra_hosts"`
		Failure     string                `json:"failure,omitempty"`
		Image       string                `json:"image,omitempty"`
		Name        string                `json:"name,omitempty"`
		Ports       []*Port               `json:"ports,omitempty"`
		Privileged  bool                  `json:"privileged,omitempty"`
		Pull        string                `json:"pull,omitempty"`
		Push        *Push                 `json:"push,omitempty"`
		Resources   *Resources            `json:"resources,omitempty"`
		Settings    map[string]*Parameter `json:"settings,omitempty"`
		Shell       string                `json:"shell,omitempty"`
		Volumes     []*VolumeMount        `json:"volumes,omitempty"`
		When        Conditions            `json:"when,omitempty"`
		WorkingDir  string                `json:"working_dir,omitempty" yaml:"working_dir"`
	}

	// Resources describes the compute resource
	// requirements.
	Resources struct {
		// Limits describes the maximum amount of compute
		// resources allowed.
		Limits *ResourceObject `json:"limits,omitempty"`

		// Requests describes the minimum amount of
		// compute resources required.
		Requests *ResourceObject `json:"requests,omitempty"`
	}

	// ResourceObject describes compute resource
	// requirements.
	ResourceObject struct {
		CPU    string    `json:"cpu"`
		Memory BytesSize `json:"memory"`
	}

	// Platform defines the target platform.
	Platform struct {
		OS      string `json:"os,omitempty"`
		Arch    string `json:"arch,omitempty"`
		Variant string `json:"variant,omitempty"`
		Version string `json:"version,omitempty"`
	}

	// Volume that can be mounted by containers.
	Volume struct {
		Name     string          `json:"name,omitempty"`
		EmptyDir *VolumeEmptyDir `json:"temp,omitempty" yaml:"temp"`
		HostPath *VolumeHostPath `json:"host,omitempty" yaml:"host"`
	}

	// VolumeDevice describes a mapping of a raw block
	// device within a container.
	VolumeDevice struct {
		Name       string `json:"name,omitempty"`
		DevicePath string `json:"path,omitempty" yaml:"path"`
	}

	// VolumeMount describes a mounting of a Volume
	// within a container.
	VolumeMount struct {
		Name      string `json:"name,omitempty"`
		MountPath string `json:"path,omitempty" yaml:"path"`
	}

	// VolumeEmptyDir mounts a temporary directory from the
	// host node's filesystem into the container. This can
	// be used as a shared scratch space.
	VolumeEmptyDir struct {
		Medium    string    `json:"medium,omitempty"`
		SizeLimit BytesSize `json:"size_limit,omitempty" yaml:"size_limit"`
	}

	// VolumeHostPath mounts a file or directory from the
	// host node's filesystem into your container.
	VolumeHostPath struct {
		Path string `json:"path,omitempty"`
	}

	// Workspace represents the pipeline workspace configuraiton.
	Workspace struct {
		Base string `json:"base,omitempty"`
		Path string `json:"path,omitempty"`
	}
)
