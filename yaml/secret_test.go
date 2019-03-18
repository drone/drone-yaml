// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package yaml

import "testing"

func TestSecretUnmarshal(t *testing.T) {
	diff, err := diff("testdata/secret.yml")
	if err != nil {
		t.Error(err)
	}
	if diff != "" {
		t.Error("Failed to parse secret")
		t.Log(diff)
	}
}

func TestSecretValidate(t *testing.T) {
	secret := new(Secret)

	secret.Data = "some-data"
	if err := secret.Validate(); err != nil {
		t.Error(err)
		return
	}

	secret.Get.Path = "secret/data/docker"
	if err := secret.Validate(); err != nil {
		t.Error(err)
		return
	}

	secret.Data = ""
	secret.Get.Path = ""
	if err := secret.Validate(); err == nil {
		t.Errorf("Expect invalid secret error")
	}
}
