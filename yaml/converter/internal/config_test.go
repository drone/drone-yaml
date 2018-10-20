package yaml

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		before, after string
	}{
		{
			before: "testdata/simple.yml",
			after:  "testdata/simple.yml.golden",
		},
		{
			before: "testdata/vault_1.yml",
			after:  "testdata/vault_1.yml.golden",
		},
		{
			before: "testdata/vault_2.yml",
			after:  "testdata/vault_2.yml.golden",
		},
		{
			before: "testdata/vault_3.yml",
			after:  "testdata/vault_3.yml.golden",
		},
	}

	for _, test := range tests {
		a, err := ioutil.ReadFile(test.before)
		if err != nil {
			t.Error(err)
			return
		}
		b, err := ioutil.ReadFile(test.after)
		if err != nil {
			t.Error(err)
			return
		}
		c, err := ConvertBytes(a)
		if err != nil {
			t.Error(err)
			return
		}
		if bytes.Equal(b, c) == false {
			t.Errorf("Unexpected yaml conversion of %s", test.before)
			t.Log(string(c))
		}
	}
}
