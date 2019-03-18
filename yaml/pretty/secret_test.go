// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package pretty

import "testing"

func TestSecret(t *testing.T) {
	ok, err := diff("testdata/secret.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}

func TestGetSecret(t *testing.T) {
	ok, err := diff("testdata/secret_get.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}
