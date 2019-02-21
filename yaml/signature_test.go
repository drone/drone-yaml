// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package yaml

import "testing"

func TestSignatureUnmarshal(t *testing.T) {
	diff, err := diff("testdata/signature.yml")
	if err != nil {
		t.Error(err)
	}
	if diff != "" {
		t.Error("Failed to parse signature")
		t.Log(diff)
	}
}

func TestSignatureValidate(t *testing.T) {
	sig := Signature{Hmac: "1234"}
	if err := sig.Validate(); err != nil {
		t.Error(err)
		return
	}

	sig.Hmac = ""
	if err := sig.Validate(); err == nil {
		t.Errorf("Expect invalid signature error")
	}
}
