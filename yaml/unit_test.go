package yaml

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestBytesSize(t *testing.T) {
	tests := []struct {
		yaml string
		size int64
		text string
	}{
		{
			yaml: "1KiB",
			size: 1024,
			text: "1KiB",
		},
		{
			yaml: "100Mi",
			size: 104857600,
			text: "100MiB",
		},
		{
			yaml: "1024",
			size: 1024,
			text: "1KiB",
		},
	}
	for _, test := range tests {
		in := []byte(test.yaml)
		out := BytesSize(0)
		err := yaml.Unmarshal(in, &out)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := int64(out), test.size; got != want {
			t.Errorf("Want byte size %d, got %d", want, got)
		}
		if got, want := out.String(), test.text; got != want {
			t.Errorf("Want byte text %s, got %s", want, got)
		}
	}
}
