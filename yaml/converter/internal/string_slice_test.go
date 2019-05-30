// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package internal

import (
	"reflect"
	"testing"

	"github.com/buildkite/yaml"
)

func TestStringSlice(t *testing.T) {
	var tests = []struct {
		yaml string
		want []string
	}{
		{
			yaml: "hello world",
			want: []string{"hello world"},
		},
		{
			yaml: "[ hello, world ]",
			want: []string{"hello", "world"},
		},
		{
			yaml: "42",
			want: []string{"42"},
		},
	}

	for _, test := range tests {
		var got StringSlice

		if err := yaml.Unmarshal([]byte(test.yaml), &got); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual([]string(got), test.want) {
			t.Errorf("Got slice %v want %v", got, test.want)
		}
	}

	var got StringSlice
	if err := yaml.Unmarshal([]byte("{}"), &got); err == nil {
		t.Errorf("Want error unmarshaling invalid string or slice value.")
	}
}
