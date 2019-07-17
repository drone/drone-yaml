// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package yaml

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		before, after, url string
	}{
		{
			before: "testdata/simple.yml",
			after:  "testdata/simple.yml.golden",
		},
		{
			before: "testdata/vault_1.yml",
			after:  "testdata/vault_1.yml.golden",
		},
		{
			before: "testdata/vault_2.yml",
			after:  "testdata/vault_2.yml.golden",
		},
		{
			before: "testdata/vault_3.yml",
			after:  "testdata/vault_3.yml.golden",
		},
		{
			before: "testdata/matrix_1.yml",
			after:  "testdata/matrix_1.yml.golden",
		},
		{
			before: "testdata/matrix_2.yml",
			after:  "testdata/matrix_2.yml.golden",
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
		c, err := Convert(a, test.url)
		if err != nil {
			t.Error(err)
			return
		}
		if bytes.Equal(b, c) == false {
			t.Errorf("Unexpected yaml conversion of %s", test.before)
			dmp := diffmatchpatch.New()
			diffs := dmp.DiffMain(string(b), string(c), false)
			t.Log(dmp.DiffCleanupSemantic(diffs))
		}
	}
}

func TestWorkspacePath(t *testing.T) {
	tests := []struct{
		a string
		b string
	}{
		{
			a: "",
			b: "src",
		},
		{
			a: "https://github.com/octocat/hello-world",
			b: "src/github.com/octocat/hello-world",
		},
		{
			a: "https://github.com:80/octocat/hello-world",
			b: "src/github.com/octocat/hello-world",
		},
		{
			a: "github.com:80/octocat/hello-world",
			b: "src",
		},
	}
	for _, test := range tests {
		if got, want := toWorkspacePath(test.a), test.b; got != want {
			t.Errorf("Want workspace path %s, got %s", want, got)
		}
	}
}