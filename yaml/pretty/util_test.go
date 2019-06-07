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

func TestQuoted(t *testing.T) {
	tests := []struct{
		before, after string
	}{
		{"", `""`},
		{"foo", "foo"},

		// special characters only quoted when followed
		// by whitespace.
		{"&foo", "&foo"},
		{"!foo", "!foo"},
		{"-foo", "-foo"},
		{":foo", ":foo"},

		{"& foo", `"& foo"`},
		{"! foo", `"! foo"`},
		{"- foo", `"- foo"`},
		{": foo", `": foo"`},

		{" & foo", `" & foo"`},
		{" ! foo", `" ! foo"`},
		{" - foo", `" - foo"`},
		{" : foo", `" : foo"`},

		// special characters only quoted when it is the
		// first character in the string.
		{",foo", `",foo"`},
		{"[foo", `"[foo"`},
		{"]foo", `"]foo"`},
		{"{foo", `"{foo"`},
		{"}foo", `"}foo"`},
		{"*foo", `"*foo"`},
		{`"foo`, `"\"foo"`},
		{`'foo`, `"'foo"`},
		{`%foo`, `"%foo"`},
		{`@foo`, `"@foo"`},
		{`|foo`, `"|foo"`},
		{`>foo`, `">foo"`},
		{`#foo`, `"#foo"`},

		{`foo:bar`, `foo:bar`},
		{`foo :bar`, `foo :bar`},
		{`foo: bar`, `"foo: bar"`},
		{`foo:`, `"foo:"`},
		{`alpine:3.8`, `alpine:3.8`}, // verify docker image names are ok

		// comments should be escaped. A comment is a pound
		// sybol preceded by a space.
		{`foo#bar`, `foo#bar`},
		{`foo #bar`, `"foo #bar"`},

		// strings with newlines and control characters
		// should be escaped
		{"foo\nbar", "\"foo\\nbar\""},
	}

	for _, test := range tests {
		buf := new(baseWriter)
		writeEncode(buf, test.before)
		a := test.after
		b := buf.String()
		if b != a {
			t.Errorf("Want %q, got %q", a, b)
		}
	}
}

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
