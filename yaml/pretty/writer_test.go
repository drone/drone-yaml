// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Community
// License that can be found in the LICENSE file.

// +build !oss

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
