// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Community
// License that can be found in the LICENSE file.

// +build !oss

package pretty

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestChunk(t *testing.T) {
	s := strings.Join(testChunk, "")
	got, want := chunk(s, 64), testChunk
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected chunk value")
		t.Log(diff)
	}
}

var testChunk = []string{
	"ZDllMjFjZDg3Zjk0ZWFjZDRhMjdhMTA1ZDQ1OTVkYTA1ODBjMTk0ZWVlZjQyNmU4",
	"N2RiNTIwZjg0NWQwYjcyYjE3MmFmZDIyYzg3NTQ1N2YyYzgxODhjYjJmNDhhOTFj",
	"ZjdhMzA0YjEzYWFlMmYxMTIwMmEyM2Q1YjQ5Yjg2ZmMK",
}

var testScalar = `>
  ZDllMjFjZDg3Zjk0ZWFjZDRhMjdhMTA1ZDQ1OTVkYTA1ODBjMTk0ZWVlZjQyNmU4
  N2RiNTIwZjg0NWQwYjcyYjE3MmFmZDIyYzg3NTQ1N2YyYzgxODhjYjJmNDhhOTFj
  ZjdhMzA0YjEzYWFlMmYxMTIwMmEyM2Q1YjQ5Yjg2ZmMK`
