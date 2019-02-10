// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Community
// License that can be found in the LICENSE file.

package yaml

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestVolume(t *testing.T) {
	var tests = []struct {
		yaml string
		want Volume
	}{
		{
			yaml: "/opt/data:/var/lib/mysql",
			want: Volume{Source: "/opt/data", Destination: "/var/lib/mysql"},
		},
		{
			yaml: "/opt/data:/var/lib/mysql:ro",
			want: Volume{Source: "/opt/data", Destination: "/var/lib/mysql", ReadOnly: true},
		},
		{
			yaml: "/opt/data:/var/lib/mysql",
			want: Volume{Source: "/opt/data", Destination: "/var/lib/mysql", ReadOnly: false},
		},
	}

	for _, test := range tests {
		got := Volume{}
		if err := yaml.Unmarshal([]byte(test.yaml), &got); err != nil {
			t.Errorf("got error unmarshaling volume %q", test.yaml)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got volume %v want %v", got, test.want)
		}
	}

	var got Volume
	if err := yaml.Unmarshal([]byte("{}"), &got); err == nil {
		t.Errorf("Want error unmarshaling invalid volume string.")
	}
}
