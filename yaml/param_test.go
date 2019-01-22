package yaml

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestParam(t *testing.T) {
	tests := []struct {
		yaml  string
		value interface{}
		from  string
	}{
		{
			yaml:  "bar",
			value: "bar",
		},
		{
			yaml: "from_secret: username",
			from: "username",
		},
	}
	for _, test := range tests {
		in := []byte(test.yaml)
		out := new(Parameter)
		err := yaml.Unmarshal(in, out)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := out.Value, test.value; got != want {
			t.Errorf("Want value %q, got %q", want, got)
		}
		if got, want := out.Secret, test.from; got != want {
			t.Errorf("Want from_secret %q, got %q", want, got)
		}
	}
}
