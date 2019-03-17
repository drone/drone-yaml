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

package pretty

import (
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

// this unit tests pretty prints a complex yaml structure
// to ensure we have common use cases covered.
func TestWriteComplexValue(t *testing.T) {
	block := map[interface{}]interface{}{}
	err := yaml.Unmarshal([]byte(testComplexValue), &block)
	if err != nil {
		t.Error(err)
		return
	}

	b := new(baseWriter)
	writeValue(b, block)
	got, want := b.String(), strings.TrimSpace(testComplexValue)
	if got != want {
		t.Errorf("Unexpected block format")
		print(got)
	}
}

var testComplexValue = `
a: b
c:
- d
- e
f:
  g: h
  i:
  - j
  - k
  - l: m
    o: p
    q:
    - r
    - s: ~
  - {}
  - []
  - ~
t: {}
u: []
v: 1
w: true
x: ~
z: "#y"
zz: "\nz\n"
"{z}": z`
