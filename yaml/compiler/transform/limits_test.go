package transform

import (
	"testing"

	"github.com/drone/drone-runtime/engine"
)

func TestWithLimits(t *testing.T) {
	step := &engine.Step{
		Metadata: engine.Metadata{
			UID:  "1",
			Name: "build",
		},
		Docker: &engine.DockerStep{},
	}
	spec := &engine.Spec{
		Steps: []*engine.Step{step},
	}
	WithLimits(1, 2)(spec)
	if got, want := step.Resources.Limits.Memory, int64(1); got != want {
		t.Errorf("Want memory limit %v, got %v", want, got)
	}
	if got, want := step.Resources.Limits.CPU, int64(2); got != want {
		t.Errorf("Want cpu limit %v, got %v", want, got)
	}
}

func TestWithNoLimits(t *testing.T) {
	step := &engine.Step{
		Metadata: engine.Metadata{
			UID:  "1",
			Name: "build",
		},
		Docker: &engine.DockerStep{},
	}
	spec := &engine.Spec{
		Steps: []*engine.Step{step},
	}
	WithLimits(0, 0)(spec)
	if step.Resources != nil {
		t.Errorf("Expect no limits applied")
	}
}
