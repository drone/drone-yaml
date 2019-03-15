// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package compiler

import (
	"testing"

	"github.com/drone/drone-runtime/engine"
	"github.com/drone/drone-yaml/yaml"
)

func TestCloneImage(t *testing.T) {
	tests := []struct {
		platform yaml.Platform
		image    string
	}{
		{
			platform: yaml.Platform{OS: "linux", Arch: "amd64"},
			image:    "drone/git",
		},
		{
			platform: yaml.Platform{OS: "linux", Arch: "arm"},
			image:    "drone/git:linux-arm",
		},
		{
			platform: yaml.Platform{OS: "linux", Arch: "arm64"},
			image:    "drone/git:linux-arm64",
		},
		{
			platform: yaml.Platform{OS: "windows", Arch: "amd64", Version: "1709"},
			image:    "drone/git:windows-1709-amd64",
		},
		{
			platform: yaml.Platform{OS: "windows", Arch: "amd64", Version: "1809"},
			image:    "drone/git:windows-1809-amd64",
		},
		{
			platform: yaml.Platform{OS: "windows", Arch: "amd64"},
			image:    "drone/git:windows-1809-amd64",
		},
		{
			platform: yaml.Platform{},
			image:    "drone/git",
		},
	}
	for _, test := range tests {
		pipeline := &yaml.Pipeline{Platform: test.platform}
		image := cloneImage(pipeline)
		if got, want := image, test.image; got != want {
			t.Errorf("Want clone image %s, got %s", want, got)
		}
	}
}

func TestSetupCloneDepth(t *testing.T) {
	// test zero depth
	src := &yaml.Pipeline{
		Clone: yaml.Clone{
			Depth: 0,
		},
	}
	dst := &engine.Step{
		Envs: map[string]string{},
	}
	setupCloneDepth(src, dst)
	if _, ok := dst.Envs["PLUGIN_DEPTH"]; ok {
		t.Errorf("Expect depth ignored when zero value")
	}

	// test non-zero depth
	src = &yaml.Pipeline{
		Clone: yaml.Clone{
			Depth: 50,
		},
	}
	dst = &engine.Step{
		Envs: map[string]string{},
	}
	setupCloneDepth(src, dst)
	if got, want := dst.Envs["PLUGIN_DEPTH"], "50"; got != want {
		t.Errorf("Expect depth %s, got %s", want, got)
	}
}

func TestSetupCloneSkipVerify(t *testing.T) {
	// test zero depth
	src := &yaml.Pipeline{
		Clone: yaml.Clone{
			SkipVerify: false,
		},
	}
	dst := &engine.Step{
		Envs: map[string]string{},
	}
	setupCloneDepth(src, dst)
	if _, ok := dst.Envs["PLUGIN_SKIP_VERIFY"]; ok {
		t.Errorf("Expect skip verify not set")
	}

	// test non-zero depth
	src = &yaml.Pipeline{
		Clone: yaml.Clone{
			SkipVerify: true,
		},
	}
	dst = &engine.Step{
		Envs: map[string]string{},
	}
	setupCloneDepth(src, dst)
	if got, want := dst.Envs["PLUGIN_SKIP_VERIFY"], "true"; got != want {
		t.Errorf("Expect skip verify %s, got %s", want, got)
	}
}
