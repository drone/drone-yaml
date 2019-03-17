// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package pretty

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/drone/drone-yaml/yaml"
)

//
//
//

func TestPrintManifest(t *testing.T) {
	ok, err := diff("testdata/manifest.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}

func diff(file string) (bool, error) {
	manifest, err := yaml.ParseFile(file)
	if err != nil {
		return false, err
	}
	golden, err := ioutil.ReadFile(file + ".golden")
	if err != nil {
		return false, err
	}

	buf := new(bytes.Buffer)
	Print(buf, manifest)

	equal := bytes.Equal(buf.Bytes(), golden)
	if !equal {
		println(">>>")
		println(buf.String())
	}
	return equal, nil
}
