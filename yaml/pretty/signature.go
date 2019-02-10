// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Community
// License that can be found in the LICENSE file.

// +build !oss

package pretty

import (
	"github.com/drone/drone-yaml/yaml"
)

// helper function pretty prints the signature resource.
func printSignature(w writer, v *yaml.Signature) {
	w.WriteString("---")
	w.WriteTagValue("version", v.Version)
	w.WriteTagValue("kind", v.Kind)
	w.WriteTagValue("hmac", v.Hmac)
	w.WriteByte('\n')
	w.WriteByte('\n')
}
