package transform

import (
	"testing"

	"github.com/drone/drone-runtime/engine"
	"github.com/google/go-cmp/cmp"
)

func TestWithSecret(t *testing.T) {
	step := &engine.Step{
		Metadata: engine.Metadata{
			UID:  "1",
			Name: "build",
		},
		Envs: map[string]string{},
	}
	spec := &engine.Spec{
		Steps: []*engine.Step{step},
	}
	secrets := map[string]string{
		"password": "correct-horse-battery-staple",
	}
	WithSecrets(secrets)(spec)

	want := []*engine.Secret{
		{
			Name: "password",
			Data: "correct-horse-battery-staple",
		},
	}
	if diff := cmp.Diff(want, spec.Secrets); diff != "" {
		t.Errorf("Unexpected secret transform")
		t.Log(diff)
	}
}

func TestWithSecretFunc(t *testing.T) {
	step := &engine.Step{
		Metadata: engine.Metadata{
			UID:  "1",
			Name: "build",
		},
		Envs: map[string]string{},
		Secrets: []*engine.SecretVar{
			{
				Name: "password",
				Env:  "PASSWORD",
			},
		},
	}
	spec := &engine.Spec{
		Steps: []*engine.Step{
			step,
			// this is a step that requests a secret
			// but should be skipped.
			{
				RunPolicy: engine.RunNever,
				Secrets: []*engine.SecretVar{
					{
						Name: "github_token",
						Env:  "GITHUB_TOKEN",
					},
				},
			},
		},
	}

	fn := func(name string) *engine.Secret {
		if name == "github_token" {
			t.Errorf("Requested secret for skipped step")
			return nil
		}
		return &engine.Secret{
			Name: "password",
			Data: "correct-horse-battery-staple",
		}
	}
	WithSecretFunc(fn)(spec)

	want := []*engine.Secret{
		{
			Name: "password",
			Data: "correct-horse-battery-staple",
		},
	}
	if diff := cmp.Diff(want, spec.Secrets); diff != "" {
		t.Errorf("Unexpected secret transform")
		t.Log(diff)
	}
}
