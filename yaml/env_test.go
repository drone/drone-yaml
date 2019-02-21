// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package yaml

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestEnv(t *testing.T) {
	tests := []struct {
		yaml  string
		value string
		name  string
		path  string
	}{
		{
			yaml:  "bar",
			value: "bar",
		},
		{
			yaml: "from_secret: username",
			name: "username",
		},
		{
			yaml: "from_secret: { name: username, path: secret/data/docker }",
			name: "username",
			path: "secret/data/docker",
		},
	}
	for _, test := range tests {
		in := []byte(test.yaml)
		out := new(Variable)
		err := yaml.Unmarshal(in, out)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := out.Value, test.value; got != want {
			t.Errorf("Want variable value %q, got %q", want, got)
		}
		if got, want := out.Secret.Name, test.name; got != want {
			t.Errorf("Want variable from_secret.name %q, got %q", want, got)
		}
		if got, want := out.Secret.Path, test.path; got != want {
			t.Errorf("Want variable from_secret.path %q, got %q", want, got)
		}
	}
}
