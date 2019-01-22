package linter

import (
	"testing"

	"github.com/drone/drone-yaml/yaml"
)

func TestManifest(t *testing.T) {
	tests := []struct {
		path    string
		trusted bool
		invalid bool
		message string
	}{
		{
			path:    "testdata/simple.yml",
			trusted: false,
			invalid: false,
		},
		{
			path:    "testdata/invalid_os.yml",
			trusted: false,
			invalid: true,
			message: "linter: unsupported os: openbsd",
		},
		{
			path:    "testdata/invalid_arch.yml",
			trusted: false,
			invalid: true,
			message: "linter: unsupported architecture: s390x",
		},
		{
			path:    "testdata/duplicate_name.yml",
			trusted: false,
			invalid: true,
			message: "linter: duplicate pipeline names",
		},
		{
			path:    "testdata/missing_dep.yml",
			trusted: false,
			invalid: true,
			message: "linter: invalid or unknown pipeline dependency",
		},
	}
	for _, test := range tests {
		t.Run(test.path, func(t *testing.T) {
			manifest, err := yaml.ParseFile(test.path)
			if err != nil {
				t.Logf("yaml: %s", test.path)
				t.Error(err)
				return
			}

			err = Manifest(manifest, test.trusted)
			if err == nil && test.invalid == true {
				t.Logf("yaml: %s", test.path)
				t.Errorf("Expect lint error")
				return
			}

			if err != nil && test.invalid == false {
				t.Logf("yaml: %s", test.path)
				t.Errorf("Expect lint error is nil, got %s", err)
				return
			}

			if err == nil {
				return
			}

			if got, want := err.Error(), test.message; got != want {
				t.Logf("yaml: %s", test.path)
				t.Errorf("Want message %q, got %q", want, got)
				return
			}
		})
	}
}
