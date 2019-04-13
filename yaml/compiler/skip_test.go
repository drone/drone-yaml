// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package compiler

import (
	"testing"

	"github.com/drone/drone-yaml/yaml"
)

func TestSkipFunc(t *testing.T) {
	tests := []struct {
		data SkipData
		when yaml.Conditions
		want bool
	}{
		//
		// test branch conditions
		//
		{
			data: SkipData{Branch: "master"},
			when: yaml.Conditions{Branch: yaml.Condition{Include: []string{"master"}}},
			want: false,
		},
		{
			data: SkipData{Branch: "master"},
			when: yaml.Conditions{Branch: yaml.Condition{Exclude: []string{"master"}}},
			want: true,
		},
		//
		// test cron conditions
		//
		{
			data: SkipData{Cron: "nightly"},
			when: yaml.Conditions{Cron: yaml.Condition{Include: []string{"nightly"}}},
			want: false,
		},
		{
			data: SkipData{Cron: "nightly"},
			when: yaml.Conditions{Cron: yaml.Condition{Exclude: []string{"nightly"}}},
			want: true,
		},
		//
		// test event conditions
		//
		{
			data: SkipData{Event: "push"},
			when: yaml.Conditions{Event: yaml.Condition{Include: []string{"push"}}},
			want: false,
		},

		{
			data: SkipData{Event: "push"},
			when: yaml.Conditions{Event: yaml.Condition{Exclude: []string{"push"}}},
			want: true,
		},
		//
		// test instance conditions
		//
		{
			data: SkipData{Instance: "drone.company.com"},
			when: yaml.Conditions{Instance: yaml.Condition{Include: []string{"drone.company.com"}}},
			want: false,
		},

		{
			data: SkipData{Instance: "drone.company.com"},
			when: yaml.Conditions{Instance: yaml.Condition{Exclude: []string{"drone.company.com"}}},
			want: true,
		},
		//
		// test ref conditions
		//
		{
			data: SkipData{Ref: "refs/heads/master"},
			when: yaml.Conditions{Ref: yaml.Condition{Include: []string{"refs/heads/*"}}},
			want: false,
		},

		{
			data: SkipData{Ref: "refs/heads/master"},
			when: yaml.Conditions{Ref: yaml.Condition{Exclude: []string{"refs/heads/*"}}},
			want: true,
		},
		//
		// test repo conditions
		//
		{
			data: SkipData{Repo: "octocat/hello-world"},
			when: yaml.Conditions{Repo: yaml.Condition{Include: []string{"octocat/hello-world"}}},
			want: false,
		},

		{
			data: SkipData{Repo: "octocat/hello-world"},
			when: yaml.Conditions{Repo: yaml.Condition{Exclude: []string{"octocat/hello-world"}}},
			want: true,
		},
		//
		// test target conditions
		//
		{
			data: SkipData{Target: "prod"},
			when: yaml.Conditions{Target: yaml.Condition{Include: []string{"prod"}}},
			want: false,
		},
		{
			data: SkipData{Target: "prod"},
			when: yaml.Conditions{Target: yaml.Condition{Exclude: []string{"prod"}}},
			want: true,
		},
	}
	for i, test := range tests {
		container := &yaml.Container{When: test.when}
		got := SkipFunc(test.data)(container)
		if got != test.want {
			t.Errorf("Want skip %v at index %d", test.want, i)
		}
	}
}
