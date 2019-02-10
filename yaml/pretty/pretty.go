// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Community
// License that can be found in the LICENSE file.

// +build !oss

package pretty

import (
	"io"

	"github.com/drone/drone-yaml/yaml"
)

// Print pretty prints the manifest.
func Print(w io.Writer, v *yaml.Manifest) {
	state := new(baseWriter)
	for _, r := range v.Resources {
		switch t := r.(type) {
		case *yaml.Cron:
			printCron(state, t)
		case *yaml.Secret:
			printSecret(state, t)
		case *yaml.Registry:
			printRegistry(state, t)
		case *yaml.Signature:
			printSignature(state, t)
		case *yaml.Pipeline:
			printPipeline(state, t)
		}
	}
	state.WriteString("...")
	state.WriteByte('\n')
	w.Write(state.Bytes())
}
