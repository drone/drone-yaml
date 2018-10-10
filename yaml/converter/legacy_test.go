package converter

import "testing"

func TestLegacy(t *testing.T) {
	tests := []struct {
		config string
		result bool
	}{
		{
			config: "\n\npipeline:\n",
			result: true,
		},
		{
			config: "\n\npipeline:   \n",
			result: true,
		},
		{
			config: "---\nkind: pipeline\n",
			result: false,
		},
	}
	for _, test := range tests {
		if got, want := IsLegacy(test.config), test.result; got != want {
			t.Errorf("Want IsLegacy %v", want)
		}
	}
	for _, test := range tests {
		b := []byte(test.config)
		if got, want := IsLegacyBytes(b), test.result; got != want {
			t.Errorf("Want IsLegacyBytes %v", want)
		}
	}
}
