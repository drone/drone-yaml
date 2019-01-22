package gitlab

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		before, after, ref string
	}{
		{
			before: "testdata/example1.yml",
			after:  "testdata/example1.yml.golden",
		},
		{
			before: "testdata/example2.yml",
			after:  "testdata/example2.yml.golden",
		},
		{
			before: "testdata/example3.yml",
			after:  "testdata/example3.yml.golden",
		},
		{
			before: "testdata/example4.yml",
			after:  "testdata/example4.yml.golden",
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
		c, err := Convert([]byte(a))
		if err != nil {
			t.Error(err)
			return
		}

		if bytes.Equal(b, c) == false {
			t.Errorf("Unexpected yaml conversion of %s", test.before)
			dmp := diffmatchpatch.New()
			diffs := dmp.DiffMain(string(b), string(c), false)
			t.Log(dmp.DiffCleanupSemantic(diffs))
		}
	}
}
