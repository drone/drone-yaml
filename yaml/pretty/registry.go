// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pretty

import (
	"github.com/drone/drone-yaml/yaml"
)

// helper function pretty prints the registry resource.
func printRegistry(w writer, v *yaml.Registry) {
	w.WriteString("---")
	w.WriteTagValue("version", v.Version)
	w.WriteTagValue("kind", v.Kind)
	w.WriteTagValue("type", v.Type)
	if v.Type == "encrypted" {
		printData(w, v.Data)
	} else {
		w.WriteTagValue("data", v.Data)
	}
	w.WriteByte('\n')
	w.WriteByte('\n')
}
