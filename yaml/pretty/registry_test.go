package pretty

import "testing"

func TestRegistry(t *testing.T) {
	ok, err := diff("testdata/registry.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}
