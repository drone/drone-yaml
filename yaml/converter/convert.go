package converter

import (
	"strings"

	"github.com/drone/drone-yaml-v1/config"
	"github.com/drone/drone-yaml-v1/yaml"
	newyaml "github.com/drone/drone-yaml/yaml"
)

// Convert from the 0.9-beta yaml to the 0.9-alpha yaml. This
// is a temporary shim and will be removed in the future.
func Convert(from *newyaml.Pipeline) *config.Config {
	to := new(config.Config)

	//
	// convert the metadata section
	//

	to.Metadata.Name = from.Name

	//
	// convert the platform section
	//

	to.Platform.OS = "linux"
	to.Platform.Arch = "amd64"
	if v := from.Platform; v != nil {
		if v.OS != "" {
			to.Platform.OS = v.OS
		}
		if v.Arch != "" {
			to.Platform.Arch = v.Arch
		}
	}

	//
	// convert the clone section
	//

	if v := from.Clone; v != nil {
		to.Clone.Depth = v.Depth
		to.Clone.Disable = v.Disable
	}

	//
	// convert the workspace section
	//

	if v := from.Workspace; v != nil {
		to.Workspace.Base = v.Base
		to.Workspace.Path = v.Path
	}

	//
	// convert the volumes section
	//

	to.Volumes = map[string]config.Volume{}
	for _, volume := range from.Volumes {
		// IGNORE host path volumes
		if volume.HostPath != nil {
			continue
		}
		// create a user-defined data-volume
		if volume.EmptyDir != nil {
			to.Volumes[volume.Name] = config.Volume{
				Driver: "local",
			}
		}
	}

	//
	// convert the trigger section
	//

	for k, v := range from.Trigger {
		if v == nil {
			continue
		}
		switch k {
		case "branch":
			to.Trigger.Branch.Include = v.Include
			to.Trigger.Branch.Exclude = v.Exclude
		case "target":
			to.Trigger.Environment.Include = v.Include
			to.Trigger.Environment.Exclude = v.Exclude
		case "event":
			to.Trigger.Event.Include = v.Include
			to.Trigger.Event.Exclude = v.Exclude
		case "instance":
			to.Trigger.Instance.Include = v.Include
			to.Trigger.Instance.Exclude = v.Exclude
		case "ref":
			to.Trigger.Ref.Include = v.Include
			to.Trigger.Ref.Exclude = v.Exclude
		case "repo":
			to.Trigger.Repo.Include = v.Include
			to.Trigger.Repo.Exclude = v.Exclude
		case "status":
			to.Trigger.Status.Include = v.Include
			to.Trigger.Status.Exclude = v.Exclude
		}
	}

	//
	// convert the depends secton
	//

	to.DependsOn = from.DependsOn

	//
	// convert the services
	//

	to.Services = map[string]*yaml.Container{}
	for _, step := range from.Services {
		to.Services[step.Name] = convertContainer(step, from)
	}

	//
	// convert the pipeline
	//

	for _, step := range from.Steps {
		to.Pipeline = append(to.Pipeline, map[string]*yaml.Container{
			step.Name: convertContainer(step, from),
		})
	}

	return to
}

func convertVargs(from map[string]*newyaml.Parameter) map[string]interface{} {
	to := map[string]interface{}{}
	for k, v := range from {
		if v.Secret != "" {
			continue
		}
		to[k] = v.Value
	}
	return to
}

func convertEnv(from map[string]*newyaml.Variable) map[string]string {
	to := map[string]string{}
	for k, v := range from {
		if v.Secret != "" {
			continue
		}
		to[k] = v.Value
	}
	return to
}

func convertConstraints(from map[string]*newyaml.Condition) yaml.Constraints {
	to := yaml.Constraints{}
	for k, v := range from {
		if v == nil {
			continue
		}
		switch k {
		case "branch":
			to.Branch.Include = v.Include
			to.Branch.Exclude = v.Exclude
		case "target":
			to.Environment.Include = v.Include
			to.Environment.Exclude = v.Exclude
		case "event":
			to.Event.Include = v.Include
			to.Event.Exclude = v.Exclude
		case "instance":
			to.Instance.Include = v.Include
			to.Instance.Exclude = v.Exclude
		case "ref":
			to.Ref.Include = v.Include
			to.Ref.Exclude = v.Exclude
		case "repo":
			to.Repo.Include = v.Include
			to.Repo.Exclude = v.Exclude
		case "status":
			to.Status.Include = v.Include
			to.Status.Exclude = v.Exclude
		}
	}
	return to
}

//
// volume conversions
//

func convertVolumeMapping(from []*newyaml.VolumeMount, pipeline *newyaml.Pipeline) []*yaml.Volume {
	var to []*yaml.Volume
	for _, v := range from {
		source := resolveVolumeSource(v.Name, pipeline)
		if source == "" {
			continue
		}
		to = append(to, &yaml.Volume{
			Source:      source,
			Destination: v.MountPath,
		})
	}
	return to
}

func resolveVolumeSource(name string, pipeline *newyaml.Pipeline) string {
	for _, volume := range pipeline.Volumes {
		if volume.Name != name {
			continue
		}
		if volume.HostPath != nil {
			return volume.HostPath.Path
		}
		if volume.EmptyDir != nil {
			return volume.Name
		}
	}
	return ""
}

//
// secret conversions
//

func convertSecretMapping(from *newyaml.Container) []*yaml.Secret {
	var to []*yaml.Secret
	for k, v := range from.Environment {
		if v.Secret != "" {
			to = append(to, &yaml.Secret{
				Source: v.Secret,
				Target: "PLUGIN" + strings.ToUpper(k),
			})
		}
	}
	for k, v := range from.Settings {
		if v.Secret != "" {
			to = append(to, &yaml.Secret{
				Source: v.Secret,
				Target: "PLUGIN" + strings.ToUpper(k),
			})
		}
	}
	return to
}

//
// container conversion
//

func convertContainer(step *newyaml.Container, pipeline *newyaml.Pipeline) *yaml.Container {
	return &yaml.Container{
		Command:     yaml.Command(step.Command),
		Commands:    yaml.StringSlice(step.Commands),
		Detached:    step.Detach,
		ErrIgnore:   step.Failure == "ignore",
		Entrypoint:  yaml.Command(step.Entrypoint),
		Environment: yaml.SliceMap{Map: convertEnv(step.Environment)},
		Image:       step.Image,
		Name:        step.Name,
		Privileged:  step.Privileged,
		Pull:        step.Pull == "always",
		Shell:       step.Shell,
		Volumes:     convertVolumeMapping(step.Volumes, pipeline),
		Secrets:     yaml.Secrets{Secrets: convertSecretMapping(step)},
		Constraints: convertConstraints(step.When),
		Vargs:       convertVargs(step.Settings),

		// TODO: devices
	}
}
