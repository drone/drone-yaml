// Copyright Jesse Haka.
// Copyright the Drone Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package yaml

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestBytesSize(t *testing.T) {
	tests := []struct {
		yaml string
		size int64
		text string
	}{
		{
			yaml: "1KiB",
			size: 1024,
			text: "1KiB",
		},
		{
			yaml: "100Mi",
			size: 104857600,
			text: "100MiB",
		},
		{
			yaml: "1024",
			size: 1024,
			text: "1KiB",
		},
	}
	for _, test := range tests {
		in := []byte(test.yaml)
		out := BytesSize(0)
		err := yaml.Unmarshal(in, &out)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := int64(out), test.size; got != want {
			t.Errorf("Want byte size %d, got %d", want, got)
		}
		if got, want := out.String(), test.text; got != want {
			t.Errorf("Want byte text %s, got %s", want, got)
		}
	}
}

func TestMilliSize(t *testing.T) {
	tests := []struct {
		yaml string
		size int64
		text string
	}{
		{
			yaml: "100m",
			size: 100,
			text: "100",
		},
		{
			yaml: "1",
			size: 1000,
			text: "1000",
		},
		{
			yaml: "0m",
			size: 0,
			text: "0",
		},
	}
	for _, test := range tests {
		in := []byte(test.yaml)
		out := MilliSize(0)
		err := yaml.Unmarshal(in, &out)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := int64(out), test.size; got != want {
			t.Errorf("Want millis %d, got %d", want, got)
		}
		if got, want := out.String(), test.text; got != want {
			t.Errorf("Want text %s, got %s", want, got)
		}
	}
}
