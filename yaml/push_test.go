package yaml

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestPush(t *testing.T) {
	tests := []struct {
		yaml  string
		image string
	}{
		{
			yaml:  "foo",
			image: "foo",
		},
		{
			yaml:  "{ image: foo }",
			image: "foo",
		},
	}
	for _, test := range tests {
		in := []byte(test.yaml)
		out := new(Push)
		err := yaml.Unmarshal(in, out)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := out.Image, test.image; got != want {
			t.Errorf("Want Image %q, got %q", want, got)
		}
	}
}

func TestPushError(t *testing.T) {
	in := []byte("[]")
	out := new(Push)
	err := yaml.Unmarshal(in, out)
	if err == nil {
		t.Errorf("Expect unmarshal error")
	}
}
