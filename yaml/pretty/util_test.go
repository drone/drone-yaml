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
