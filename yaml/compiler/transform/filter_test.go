// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package transform

import (
	"testing"

	"github.com/drone/drone-runtime/engine"
)

func TestInclude(t *testing.T) {
	step1 := &engine.Step{
		Metadata:  engine.Metadata{Name: "clone"},
		RunPolicy: engine.RunOnSuccess,
	}
	step2 := &engine.Step{
		Metadata:  engine.Metadata{Name: "build"},
		RunPolicy: engine.RunOnSuccess,
	}
	step3 := &engine.Step{
		Metadata:  engine.Metadata{Name: "test"},
		RunPolicy: engine.RunOnSuccess,
	}
	spec := &engine.Spec{
		Metadata: engine.Metadata{},
		Steps:    []*engine.Step{step1, step2, step3},
	}
	Include([]string{"test"})(spec)
	if got, want := step1.RunPolicy, engine.RunOnSuccess; got != want {
		t.Errorf("Want run policy %s got %s", want, got)
	}
	if got, want := step2.RunPolicy, engine.RunNever; got != want {
		t.Errorf("Want run policy %s got %s", want, got)
	}
	if got, want := step3.RunPolicy, engine.RunOnSuccess; got != want {
		t.Errorf("Want run policy %s got %s", want, got)
	}
}

func TestInclude_Empty(t *testing.T) {
	step := &engine.Step{
		Metadata:  engine.Metadata{},
		RunPolicy: engine.RunOnSuccess,
	}
	spec := &engine.Spec{
		Metadata: engine.Metadata{},
		Steps:    []*engine.Step{step},
	}
	Include(nil)(spec)
	if got, want := step.RunPolicy, engine.RunOnSuccess; got != want {
		t.Errorf("Want run policy %s got %s", want, got)
	}
}

func TestExclude(t *testing.T) {
	step1 := &engine.Step{
		Metadata:  engine.Metadata{Name: "clone"},
		RunPolicy: engine.RunOnSuccess,
	}
	step2 := &engine.Step{
		Metadata:  engine.Metadata{Name: "build"},
		RunPolicy: engine.RunOnSuccess,
	}
	step3 := &engine.Step{
		Metadata:  engine.Metadata{Name: "test"},
		RunPolicy: engine.RunOnSuccess,
	}
	spec := &engine.Spec{
		Metadata: engine.Metadata{},
		Steps:    []*engine.Step{step1, step2, step3},
	}
	Exclude([]string{"test"})(spec)
	if got, want := step1.RunPolicy, engine.RunOnSuccess; got != want {
		t.Errorf("Want run policy %s got %s", want, got)
	}
	if got, want := step2.RunPolicy, engine.RunOnSuccess; got != want {
		t.Errorf("Want run policy %s got %s", want, got)
	}
	if got, want := step3.RunPolicy, engine.RunNever; got != want {
		t.Errorf("Want run policy %s got %s", want, got)
	}
}

func TestExclude_Empty(t *testing.T) {
	step := &engine.Step{
		Metadata:  engine.Metadata{},
		RunPolicy: engine.RunOnSuccess,
	}
	spec := &engine.Spec{
		Metadata: engine.Metadata{},
		Steps:    []*engine.Step{step},
	}
	Exclude(nil)(spec)
	if got, want := step.RunPolicy, engine.RunOnSuccess; got != want {
		t.Errorf("Want run policy %s got %s", want, got)
	}
}

func TestResumeAt(t *testing.T) {
	step1 := &engine.Step{
		Metadata:  engine.Metadata{Name: "clone"},
		RunPolicy: engine.RunOnSuccess,
	}
	step2 := &engine.Step{
		Metadata:  engine.Metadata{Name: "build"},
		RunPolicy: engine.RunOnSuccess,
	}
	step3 := &engine.Step{
		Metadata:  engine.Metadata{Name: "test"},
		RunPolicy: engine.RunOnSuccess,
	}
	spec := &engine.Spec{
		Metadata: engine.Metadata{},
		Steps:    []*engine.Step{step1, step2, step3},
	}
	ResumeAt("test")(spec)
	if got, want := step1.RunPolicy, engine.RunOnSuccess; got != want {
		t.Errorf("Want run policy %s got %s", want, got)
	}
	if got, want := step2.RunPolicy, engine.RunNever; got != want {
		t.Errorf("Want run policy %s got %s", want, got)
	}
	if got, want := step3.RunPolicy, engine.RunOnSuccess; got != want {
		t.Errorf("Want run policy %s got %s", want, got)
	}
}

func TestResume_Empty(t *testing.T) {
	step := &engine.Step{
		Metadata:  engine.Metadata{},
		RunPolicy: engine.RunOnSuccess,
	}
	spec := &engine.Spec{
		Metadata: engine.Metadata{},
		Steps:    []*engine.Step{step},
	}
	ResumeAt("")(spec)
	if got, want := step.RunPolicy, engine.RunOnSuccess; got != want {
		t.Errorf("Want run policy %s got %s", want, got)
	}
}
