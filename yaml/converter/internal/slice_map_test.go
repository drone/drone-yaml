package yaml

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestMapSlice(t *testing.T) {
	var tests = []struct {
		yaml string
		want map[string]string
	}{
		{
			yaml: "[ foo=bar, baz=qux ]",
			want: map[string]string{"foo": "bar", "baz": "qux"},
		},
		{
			yaml: "{ foo: bar, baz: qux }",
			want: map[string]string{"foo": "bar", "baz": "qux"},
		},
	}

	for _, test := range tests {
		var got SliceMap

		if err := yaml.Unmarshal([]byte(test.yaml), &got); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(got.Map, test.want) {
			t.Errorf("Got map %v want %v", got, test.want)
		}
	}

	var got SliceMap
	if err := yaml.Unmarshal([]byte("1"), &got); err == nil {
		t.Errorf("Want error unmarshaling invalid map value.")
	}
}
