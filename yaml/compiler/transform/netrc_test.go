package transform

import (
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/drone/drone-runtime/engine"
	"github.com/google/go-cmp/cmp"
)

var ignoreMetadata = cmpopts.IgnoreFields(
	engine.Metadata{}, "UID")

func TestWithNetrc(t *testing.T) {
	step := &engine.Step{
		Metadata: engine.Metadata{
			UID:  "1",
			Name: "build",
		},
	}
	spec := &engine.Spec{
		Metadata: engine.Metadata{
			UID:       "acdj0yjqv7uh5hidveg0ggr42x8oj67b",
			Namespace: "pivqfthg1c9hy83ylht1sxx4nygjc7tk",
		},
		Steps: []*engine.Step{step},
	}
	WithNetrc("@machine", "@username", "@password")(spec)
	if len(step.Files) == 0 {
		t.Errorf("File mount not added to step")
		return
	}
	if len(spec.Files) == 0 {
		t.Errorf("File not declared in spec")
		return
	}
	file := &engine.File{
		Metadata: engine.Metadata{
			Name:      ".netrc",
			Namespace: "pivqfthg1c9hy83ylht1sxx4nygjc7tk",
		},
		Data: []byte("machine @machine login @username password @password"),
	}
	if diff := cmp.Diff(file, spec.Files[0], ignoreMetadata); diff != "" {
		t.Errorf("Unexpected file declaration")
		t.Log(diff)
	}

	fileMount := &engine.FileMount{Name: ".netrc", Path: "/root/.netrc", Mode: 0600}
	if diff := cmp.Diff(fileMount, step.Files[0], ignoreMetadata); diff != "" {
		t.Errorf("Unexpected file mount")
		t.Log(diff)
	}
}

func TestWithEmptyNetrc(t *testing.T) {
	step := &engine.Step{
		Metadata: engine.Metadata{
			UID:  "1",
			Name: "build",
		},
	}
	spec := &engine.Spec{
		Steps: []*engine.Step{step},
	}
	WithNetrc("@machine", "", "")(spec)
	if len(spec.Files) != 0 {
		t.Errorf("Unexpected file declaration")
	}
	if len(step.Files) != 0 {
		t.Errorf("Unexpected file mount")
	}
}
