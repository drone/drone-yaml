// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Community
// License that can be found in the LICENSE file.

package yaml

import "testing"

func TestPipelineUnmarshal(t *testing.T) {
	diff, err := diff("testdata/pipeline.yml")
	if err != nil {
		t.Error(err)
	}
	if diff != "" {
		t.Error("Failed to parse pipeline")
		t.Log(diff)
	}
}
