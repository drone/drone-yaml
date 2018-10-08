package transform

import (
	"testing"

	"github.com/drone/drone-runtime/engine"
)

func TestWithVolumes(t *testing.T) {
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
		Docker: &engine.DockerConfig{},
		Steps:  []*engine.Step{step},
	}
	vols := map[string]string{"/path/on/host": "/path/in/container"}
	WithVolumes(vols)(spec)

	if len(step.Volumes) == 0 {
		t.Error("Expected volume added to container")
	}
	if got, want := step.Volumes[0].Path, "/path/in/container"; got != want {
		t.Errorf("Want mount path %s, got %s", want, got)
	}
	if len(spec.Docker.Volumes) == 0 {
		t.Error("Expected volume added to spec")
	}
	if got, want := spec.Docker.Volumes[0].HostPath.Path, "/path/on/host"; got != want {
		t.Errorf("Want host mount path %s, got %s", want, got)
	}
}
