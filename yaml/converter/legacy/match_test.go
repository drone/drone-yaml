// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package legacy

import "testing"

func TestMatch(t *testing.T) {
	tests := []struct {
		config string
		result bool
	}{
		{
			config: "pipeline:\n  build:\n    image: golang:1.11",
			result: true,
		},
		{
			config: "\n\npipeline:\n",
			result: true,
		},
		{
			config: "\n\npipeline:   \n",
			result: true,
		},
		{
			config: "---\nkind: pipeline\n",
			result: false,
		},
	}
	for i, test := range tests {
		b := []byte(test.config)
		if got, want := Match(b), test.result; got != want {
			t.Errorf("Want IsLegacyBytes %v at index %d,", want, i)
		}
	}
}
