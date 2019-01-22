package transform

import (
	"testing"

	"github.com/drone/drone-runtime/engine"
	"github.com/google/go-cmp/cmp"
)

func TestWithNetwork(t *testing.T) {
	step := &engine.Step{
		Metadata: engine.Metadata{
			UID:  "1",
			Name: "build",
		},
		Docker: &engine.DockerStep{
			Networks: nil,
		},
	}
	spec := &engine.Spec{
		Steps: []*engine.Step{step},
	}
	nets := []string{"a", "b"}
	WithNetworks(nets)(spec)
	if diff := cmp.Diff(nets, step.Docker.Networks); diff != "" {
		t.Errorf("Unexpected transform")
		t.Log(diff)
	}
}
