// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package yaml

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestParam(t *testing.T) {
	tests := []struct {
		yaml  string
		value interface{}
		name  string
		path  string
	}{
		{
			yaml:  "bar",
			value: "bar",
			name:  "",
			path:  "",
		},
		{
			yaml:  "[ bar ]",
			value: []interface{}{"bar"},
			name:  "",
			path:  "",
		},
		{
			yaml:  "from_secret: username",
			value: nil,
			name:  "username",
			path:  "",
		},
		{
			yaml:  "from_secret: { path: secret/data/docker, name: username }",
			value: nil,
			name:  "username",
			path:  "secret/data/docker",
		},
	}
	for _, test := range tests {
		in := []byte(test.yaml)
		out := new(Parameter)
		err := yaml.Unmarshal(in, out)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := out.Value, test.value; !reflect.DeepEqual(got, want) {
			t.Errorf("Want value %q of type %T, got %q of type %T", want, want, got, got)
		}
		if got, want := out.Secret.Name, test.name; got != want {
			t.Errorf("Want from_secret.name %q, got %q", want, got)
		}
		if got, want := out.Secret.Path, test.path; got != want {
			t.Errorf("Want from_secret.path %q, got %q", want, got)
		}
	}
}
