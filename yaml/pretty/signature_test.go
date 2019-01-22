package pretty

import "testing"

func TestSignature(t *testing.T) {
	ok, err := diff("testdata/signature.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}
