// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package transform

import (
	"os"
	"testing"

	"github.com/drone/drone-runtime/engine"
)

func TestWithProxy(t *testing.T) {
	var (
		noProxy    = getenv("no_proxy")
		httpProxy  = getenv("https_proxy")
		httpsProxy = getenv("https_proxy")
	)
	defer func() {
		os.Setenv("no_proxy", noProxy)
		os.Setenv("NO_PROXY", noProxy)
		os.Setenv("http_proxy", httpProxy)
		os.Setenv("HTTP_PROXY", httpProxy)
		os.Setenv("HTTPS_PROXY", httpsProxy)
		os.Setenv("https_proxy", httpsProxy)
	}()

	testdata := map[string]string{
		"NO_PROXY":    "http://dummy.no.proxy",
		"http_proxy":  "http://dummy.http.proxy",
		"https_proxy": "http://dummy.https.proxy",
	}

	for k, v := range testdata {
		os.Setenv(k, v)
	}

	step := &engine.Step{
		Metadata: engine.Metadata{
			UID:  "1",
			Name: "build",
		},
		Envs: map[string]string{},
	}
	spec := &engine.Spec{
		Steps: []*engine.Step{step},
	}
	WithProxy()(spec)
	for k, v := range testdata {
		step := spec.Steps[0]
		if step.Envs[k] != v {
			t.Errorf("Expect proxy varaible %s=%q, got %q", k, v, step.Envs[k])
		}
	}
}
