// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package transform

import (
	"testing"

	"github.com/drone/drone-runtime/engine"
	"github.com/google/go-cmp/cmp"
)

func TestWithLabels(t *testing.T) {
	volume := &engine.Volume{
		Metadata: engine.Metadata{
			Labels: map[string]string{},
		},
	}
	step := &engine.Step{
		Metadata: engine.Metadata{
			Labels: map[string]string{},
		},
		Envs: map[string]string{},
	}
	spec := &engine.Spec{
		Metadata: engine.Metadata{
			Labels: map[string]string{},
		},
		Steps: []*engine.Step{step},
		Docker: &engine.DockerConfig{
			Volumes: []*engine.Volume{volume},
		},
	}
	labels := map[string]string{
		"io.drone.build.number": "1",
		"io.drone.build.event":  "push",
	}
	WithLables(labels)(spec)
	if diff := cmp.Diff(labels, spec.Metadata.Labels); diff != "" {
		t.Errorf("Unexpected spec labels")
		t.Log(diff)
	}
	if diff := cmp.Diff(labels, step.Metadata.Labels); diff != "" {
		t.Errorf("Unexpected step labels")
		t.Log(diff)
	}
	if diff := cmp.Diff(labels, volume.Metadata.Labels); diff != "" {
		t.Errorf("Unexpected volume labels")
		t.Log(diff)
	}
}
