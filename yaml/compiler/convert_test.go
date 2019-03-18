// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package compiler

import (
	"testing"

	"github.com/drone/drone-runtime/engine"
	"github.com/drone/drone-yaml/yaml"
	"github.com/google/go-cmp/cmp"
)

func Test_toIgnoreErr(t *testing.T) {
	tests := []struct {
		mode string
		want bool
	}{
		{"Ignore", true},
		{"ignore", true},
		{"fail", false},
	}
	for _, test := range tests {
		from := &yaml.Container{Failure: test.mode}
		if toIgnoreErr(from) != test.want {
			t.Errorf("unexpected ignore error for %s", test.mode)
		}
	}
}

func Test_toPullPolicy(t *testing.T) {
	tests := []struct {
		mode string
		want engine.PullPolicy
	}{
		{"", engine.PullDefault},
		{"always", engine.PullAlways},
		{"if-not-exists", engine.PullIfNotExists},
		{"never", engine.PullNever},
	}
	for _, test := range tests {
		from := &yaml.Container{Pull: test.mode}
		if toPullPolicy(from) != test.want {
			t.Errorf("unexpected pull policy for %s", test.mode)
		}
	}
}

func Test_toRunPolicy(t *testing.T) {
	tests := []struct {
		cond yaml.Condition
		want engine.RunPolicy
	}{
		{yaml.Condition{}, engine.RunOnSuccess},
		{yaml.Condition{Include: []string{"success"}}, engine.RunOnSuccess},
		{yaml.Condition{Include: []string{"failure"}}, engine.RunOnFailure},
		{yaml.Condition{Include: []string{"success", "failure"}}, engine.RunAlways},
		{yaml.Condition{Exclude: []string{"success", "failure"}}, engine.RunNever},
	}
	for _, test := range tests {
		from := &yaml.Container{When: yaml.Conditions{Status: test.cond}}
		if toRunPolicy(from) != test.want {
			t.Errorf("unexpected pull policy for incude: %s, exclude: %s", test.cond.Include, test.cond.Exclude)
		}
	}
}

func Test_toPorts(t *testing.T) {
	from := &yaml.Container{
		Ports: []*yaml.Port{
			{
				Port:     80,
				Host:     8080,
				Protocol: "TCP",
			},
			{
				Port:     80,
				Host:     0,
				Protocol: "",
			},
		},
	}
	a := toPorts(from)
	b := []*engine.Port{
		{
			Port:     80,
			Host:     8080,
			Protocol: "TCP",
		},
		{
			Port:     80,
			Host:     0,
			Protocol: "",
		},
	}
	if diff := cmp.Diff(a, b); diff != "" {
		t.Errorf("Unexpected port conversion")
		t.Log(diff)
	}
}

func Test_toResources(t *testing.T) {
	from := &yaml.Container{
		Resources: nil,
	}
	if toResources(from) != nil {
		t.Errorf("Expected nil resources")
	}

	// test what happens when limits are defined
	// but reservations are nil.

	from = &yaml.Container{
		Resources: &yaml.Resources{
			Limits: &yaml.ResourceObject{
				Memory: yaml.BytesSize(1000),
			},
		},
	}
	a := toResources(from)
	b := &engine.Resources{
		Limits: &engine.ResourceObject{
			Memory: 1000,
		},
	}
	if diff := cmp.Diff(a, b); diff != "" {
		t.Errorf("Unexpected resource conversion")
		t.Log(diff)
	}

	// test what happens when reservation and limits
	// are both provided.

	from = &yaml.Container{
		Resources: &yaml.Resources{
			Limits: &yaml.ResourceObject{
				Memory: yaml.BytesSize(1000),
				CPU:    4,
			},
			Requests: &yaml.ResourceObject{
				Memory: yaml.BytesSize(2000),
				CPU:    0.1,
			},
		},
	}
	a = toResources(from)
	b = &engine.Resources{
		Limits: &engine.ResourceObject{
			Memory: 1000,
			CPU:    4000,
		},
		Requests: &engine.ResourceObject{
			Memory: 2000,
			CPU:    100,
		},
	}
	if diff := cmp.Diff(a, b); diff != "" {
		t.Errorf("Unexpected resource conversion")
		t.Log(diff)
	}
}
