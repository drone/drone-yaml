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

package yaml

import "errors"

type (
	// Secret is a resource that provides encrypted data
	// and pointers to external data (i.e. from vault).
	Secret struct {
		Version string `json:"version,omitempty"`
		Kind    string `json:"kind,omitempty"`
		Type    string `json:"type,omitempty"`

		Data     map[string]string       `json:"data,omitempty"`
		External map[string]ExternalData `json:"external_data,omitempty" yaml:"external_data"`
	}

	// ExternalData defines the path and name of external
	// data located in an external or remote storage system.
	ExternalData struct {
		Path string `json:"path,omitempty"`
		Name string `json:"name,omitempty"`
	}
)

// GetVersion returns the resource version.
func (s *Secret) GetVersion() string { return s.Version }

// GetKind returns the resource kind.
func (s *Secret) GetKind() string { return s.Kind }

// Validate returns an error if the secret is invalid.
func (s *Secret) Validate() error {
	if len(s.Data) == 0 && len(s.External) == 0 {
		return errors.New("yaml: invalid secret resource")
	}
	return nil
}
