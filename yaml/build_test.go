// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Community
// License that can be found in the LICENSE file.

package yaml

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		yaml  string
		image string
	}{
		{
			yaml:  "bar",
			image: "bar",
		},
		{
			yaml:  "{ image: foo }",
			image: "foo",
		},
	}
	for _, test := range tests {
		in := []byte(test.yaml)
		out := new(Build)
		err := yaml.Unmarshal(in, out)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := out.Image, test.image; got != want {
			t.Errorf("Want image %q, got %q", want, got)
		}
	}
}

func TestBuildError(t *testing.T) {
	in := []byte("[]")
	out := new(Build)
	err := yaml.Unmarshal(in, out)
	if err == nil {
		t.Errorf("Expect unmarshal error")
	}
}
