package compiler

import (
	"testing"

	"github.com/drone/drone-runtime/engine"
	"github.com/drone/drone-yaml/yaml"
)

func TestSetupWorkspace(t *testing.T) {
	tests := []struct {
		path string
		src  *yaml.Container
		dst  *engine.Step
		want string
	}{
		{
			path: "/drone/src",
			src:  &yaml.Container{},
			dst:  &engine.Step{},
			want: "/drone/src",
		},
		// do not override the user-defined working dir.
		{
			path: "/drone/src",
			src:  &yaml.Container{},
			dst:  &engine.Step{WorkingDir: "/foo"},
			want: "/foo",
		},
		// do not override the default working directory
		// for service containers with no commands.
		{
			path: "/drone/src",
			src:  &yaml.Container{},
			dst:  &engine.Step{Detach: true},
			want: "",
		},
		// overrides the default working directory
		// for service containers with commands.
		{
			path: "/drone/src",
			src:  &yaml.Container{Commands: []string{"whoami"}},
			dst:  &engine.Step{Detach: true},
			want: "/drone/src",
		},
	}
	for _, test := range tests {
		setupWorkingDir(test.src, test.dst, test.path)
		if got, want := test.dst.WorkingDir, test.want; got != want {
			t.Errorf("Want working_dir %s, got %s", want, got)
		}
	}
}

func TestToWindows(t *testing.T) {
	got := toWindowsDrive("/go/src/github.com/octocat/hello-world")
	want := "c:\\go\\src\\github.com\\octocat\\hello-world"
	if got != want {
		t.Errorf("Want windows drive %q, got %q", want, got)
	}
}

func TestCreateWorkspace(t *testing.T) {
	tests := []struct {
		from *yaml.Pipeline
		base string
		path string
		full string
	}{
		{
			from: &yaml.Pipeline{},
			base: "/drone",
			path: "src",
			full: "/drone/src",
		},
		{
			from: &yaml.Pipeline{
				Workspace: yaml.Workspace{
					Base: "/foo",
					Path: "/bar",
				},
			},
			base: "/foo",
			path: "/bar",
			full: "/foo/bar",
		},
		{
			from: &yaml.Pipeline{
				Platform: yaml.Platform{
					OS: "windows",
				},
			},
			base: "c:\\drone",
			path: "src",
			full: "c:\\drone\\src",
		},
	}
	for _, test := range tests {
		base, path, full := createWorkspace(test.from)
		if got, want := test.base, base; got != want {
			t.Errorf("Want workspace base %s, got %s", want, got)
		}
		if got, want := test.path, path; got != want {
			t.Errorf("Want workspace path %s, got %s", want, got)
		}
		if got, want := test.full, full; got != want {
			t.Errorf("Want workspace %s, got %s", want, got)
		}
	}
}
