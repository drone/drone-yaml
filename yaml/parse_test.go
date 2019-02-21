// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package yaml

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseRaw(t *testing.T) {
	tests := []struct {
		data string
		want []*RawResource
	}{
		//
		// empty document returns nil resources.
		//
		{
			data: "",
			want: nil,
		},
		//
		// single document files.
		//
		{
			data: "kind: pipeline\nfoo: bar",
			want: []*RawResource{
				{Kind: "pipeline", Data: []byte("kind: pipeline\nfoo: bar\n")},
			},
		},
		{
			data: "kind: pipeline\nfoo: bar\n",
			want: []*RawResource{
				{Kind: "pipeline", Data: []byte("kind: pipeline\nfoo: bar\n")},
			},
		},
		{
			data: "---\nkind: pipeline\nfoo: bar\n...",
			want: []*RawResource{
				{Kind: "pipeline", Data: []byte("kind: pipeline\nfoo: bar\n")},
			},
		},
		{
			data: "---\nkind: pipeline\nfoo: bar\n...\n",
			want: []*RawResource{
				{Kind: "pipeline", Data: []byte("kind: pipeline\nfoo: bar\n")},
			},
		},
		//
		// multi-document files.
		//
		{
			data: "kind: a\nb: c\n---\nkind: d\ne: f",
			want: []*RawResource{
				{Kind: "a", Data: []byte("kind: a\nb: c\n")},
				{Kind: "d", Data: []byte("kind: d\ne: f\n")},
			},
		},
		{
			data: "---\nkind: a\nb: c\n---\nkind: d\ne: f\n",
			want: []*RawResource{
				{Kind: "a", Data: []byte("kind: a\nb: c\n")},
				{Kind: "d", Data: []byte("kind: d\ne: f\n")},
			},
		},
		{
			data: "---\nkind: a\nb: c\n---\nkind: d\ne: f\n...",
			want: []*RawResource{
				{Kind: "a", Data: []byte("kind: a\nb: c\n")},
				{Kind: "d", Data: []byte("kind: d\ne: f\n")},
			},
		},
		{
			data: "---\nkind: a\nb: c\n---\nkind: d\ne: f\n...\n",
			want: []*RawResource{
				{Kind: "a", Data: []byte("kind: a\nb: c\n")},
				{Kind: "d", Data: []byte("kind: d\ne: f\n")},
			},
		},
	}

	for i, test := range tests {
		name := fmt.Sprint(i)
		t.Run(name, func(t *testing.T) {
			got, err := ParseRawString(test.data)
			if err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Fail()
				t.Log(diff)
			}
		})
	}
}
