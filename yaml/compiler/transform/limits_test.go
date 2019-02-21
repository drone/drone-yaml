// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

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
	if got, want := step.Resources.Limits.CPU, int64(2000); got != want {
		t.Errorf("Want cpu limit %v, got %v", want, got)
	}
}

func TestWithMemory(t *testing.T) {
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
	WithLimits(1, 0)(spec)
	if got, want := step.Resources.Limits.Memory, int64(1); got != want {
		t.Errorf("Want memory limit %v, got %v", want, got)
	}
	if got, want := step.Resources.Limits.CPU, int64(0); got != want {
		t.Errorf("Want cpu limit %v, got %v", want, got)
	}
}

func TestWithCPU(t *testing.T) {
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
	WithLimits(0, 3)(spec)
	if got, want := step.Resources.Limits.Memory, int64(0); got != want {
		t.Errorf("Want memory limit %v, got %v", want, got)
	}
	if got, want := step.Resources.Limits.CPU, int64(3000); got != want {
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
