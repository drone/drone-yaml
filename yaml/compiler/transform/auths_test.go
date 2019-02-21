// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package transform

import (
	"testing"

	"github.com/drone/drone-runtime/engine"
	"github.com/google/go-cmp/cmp"
)

func TestWithAuths(t *testing.T) {
	spec := &engine.Spec{
		Steps:  []*engine.Step{},
		Docker: &engine.DockerConfig{},
	}
	auths := []*engine.DockerAuth{
		{
			Address:  "docker.io",
			Username: "octocat",
			Password: "correct-horse-battery-staple",
		},
	}
	WithAuths(auths)(spec)
	if diff := cmp.Diff(auths, spec.Docker.Auths); diff != "" {
		t.Errorf("Unexpected auths transform")
		t.Log(diff)
	}
}

func TestWithAuthsFunc(t *testing.T) {
	spec := &engine.Spec{
		Steps:  []*engine.Step{},
		Docker: &engine.DockerConfig{},
	}
	auths := []*engine.DockerAuth{
		{
			Address:  "docker.io",
			Username: "octocat",
			Password: "correct-horse-battery-staple",
		},
	}
	fn := func() []*engine.DockerAuth {
		return auths
	}
	WithAuthsFunc(fn)(spec)
	if diff := cmp.Diff(auths, spec.Docker.Auths); diff != "" {
		t.Errorf("Unexpected auths transform")
		t.Log(diff)
	}
}
