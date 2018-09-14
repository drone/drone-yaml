package yaml

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestManifestUnmarshal(t *testing.T) {
	diff, err := diff("testdata/manifest.yml")
	if err != nil {
		t.Error(err)
	}
	if diff != "" {
		t.Error("Failed to parse manifest with multiple entries")
		t.Log(diff)
	}
}

func diff(file string) (string, error) {
	a, err := ParseFile(file)
	if err != nil {
		return "", err
	}
	// enc := json.NewEncoder(os.Stdout)
	// enc.SetIndent("", "  ")
	// enc.Encode(a)
	d, err := ioutil.ReadFile(file + ".golden")
	if err != nil {
		return "", err
	}
	b := new(Manifest)
	err = json.Unmarshal(d, b)
	if err != nil {
		return "", err
	}
	return cmp.Diff(a, b), nil
}
