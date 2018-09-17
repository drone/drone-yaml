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

	to.Platform.OS = from.Platform.OS
	to.Platform.Arch = from.Platform.Arch
	if to.Platform.OS == "" {
		to.Platform.OS = "linux"
	}
	if to.Platform.Arch == "" {
		to.Platform.Arch = "amd64"
	}
	to.Platform.Name = to.Platform.OS + "/" + to.Platform.Arch

	//
	// convert the clone section
	//

	to.Clone.Depth = from.Clone.Depth
	to.Clone.Disable = from.Clone.Disable

	//
	// convert the workspace section
	//

	to.Workspace.Base = from.Workspace.Base
	to.Workspace.Path = from.Workspace.Path

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

	to.Trigger = convertConstraints(from.Trigger)

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

func convertConstraints(from newyaml.Conditions) yaml.Constraints {
	to := yaml.Constraints{}
	to.Branch.Include = from.Branch.Include
	to.Branch.Exclude = from.Branch.Exclude
	to.Environment.Include = from.Target.Include
	to.Environment.Exclude = from.Target.Exclude
	to.Event.Include = from.Event.Include
	to.Event.Exclude = from.Event.Exclude
	to.Instance.Include = from.Instance.Include
	to.Instance.Exclude = from.Instance.Exclude
	to.Ref.Include = from.Ref.Include
	to.Ref.Exclude = from.Ref.Exclude
	to.Repo.Include = from.Repo.Include
	to.Repo.Exclude = from.Repo.Exclude
	to.Status.Include = from.Status.Include
	to.Status.Exclude = from.Status.Exclude
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
				Target: "PLUGIN_" + strings.ToUpper(k),
			})
			to = append(to, &yaml.Secret{
				Source: v.Secret,
				Target: strings.ToUpper(k),
			})
		}
	}
	for k, v := range from.Settings {
		if v.Secret != "" {
			to = append(to, &yaml.Secret{
				Source: v.Secret,
				Target: "PLUGIN_" + strings.ToUpper(k),
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
