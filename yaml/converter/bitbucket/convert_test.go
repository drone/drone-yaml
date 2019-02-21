// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package bitbucket

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		before, after, ref string
	}{
		{
			before: "testdata/sample1.yaml",
			after:  "testdata/sample1.yaml.golden",
			ref:    "refs/heads/master",
		},
		{
			before: "testdata/sample2.yaml",
			after:  "testdata/sample2.yaml.golden",
			ref:    "refs/heads/feature/foo",
		},
	}

	for _, test := range tests {
		a, err := ioutil.ReadFile(test.before)
		if err != nil {
			t.Error(err)
			return
		}
		b, err := ioutil.ReadFile(test.after)
		if err != nil {
			t.Error(err)
			return
		}
		c, err := Convert([]byte(a), test.ref)
		if err != nil {
			t.Error(err)
			return
		}
		if bytes.Equal(b, c) == false {
			t.Errorf("Unexpected yaml conversion of %s", test.before)
			t.Log(string(c))
		}
	}
}
