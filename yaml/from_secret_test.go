// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package yaml

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestFromSecret(t *testing.T) {
	tests := []struct {
		yaml string
		name string
		path string
	}{
		{
			yaml: "bar",
			name: "bar",
			path: "",
		},
		{
			yaml: "{ name: bar }",
			name: "bar",
			path: "",
		},
		{
			yaml: "{ name: bar, path: foo }",
			name: "bar",
			path: "foo",
		},
	}
	for _, test := range tests {
		in := []byte(test.yaml)
		out := new(FromSecret)
		err := yaml.Unmarshal(in, out)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := out.Name, test.name; got != want {
			t.Errorf("Want from_secret name %q, got %q", want, got)
		}
		if got, want := out.Path, test.path; got != want {
			t.Errorf("Want from_secret path %q, got %q", want, got)
		}
	}
}
