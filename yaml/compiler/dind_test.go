package compiler

import (
	"testing"

	"github.com/drone/drone-yaml/yaml"
)

func TestDind(t *testing.T) {
	tests := []struct {
		container  *yaml.Container
		privileged bool
	}{
		{
			container:  &yaml.Container{Image: "plugins/docker"},
			privileged: true,
		},
		{
			container:  &yaml.Container{Image: "plugins/docker:latest"},
			privileged: true,
		},
		{
			container:  &yaml.Container{Image: "plugins/docker:1"},
			privileged: true,
		},
		// no match
		{
			container:  &yaml.Container{Image: "golang"},
			privileged: false,
		},
		// dind containers cannot set entrypoint or command
		{
			container: &yaml.Container{
				Image:   "plugins/docker",
				Command: []string{"docker run ..."},
			},
			privileged: false,
		},
		{
			container: &yaml.Container{
				Image:      "plugins/docker",
				Entrypoint: []string{"docker run ..."},
			},
			privileged: false,
		},
		// dind containers cannot set commands
		{
			container: &yaml.Container{
				Image:    "plugins/docker",
				Commands: []string{"docker run ..."},
			},
			privileged: false,
		},
	}
	for i, test := range tests {
		images := []string{"plugins/docker", "plugins/ecr"}
		privileged := DindFunc(images)(test.container)
		if privileged != test.privileged {
			t.Errorf("Want privileged %v at index %d", test.privileged, i)
		}
	}

}
