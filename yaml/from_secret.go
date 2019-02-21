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

type (
	// FromSecret defines a secret sources from an external value.
	// This could be a named secret stored in the Drone database,
	// or a secret stored at an external path in a key value
	// system such as Vault.
	FromSecret struct {
		Name string `json:"name"`
		Path string `json:"path"`
	}

	// fromSecret is a tempoary type used to unmarshal
	// from_secret which may be short form vs long form.
	fromSecret struct {
		Name string `json:"name"`
		Path string `json:"path"`
	}
)

// UnmarshalYAML implements yaml unmarshalling.
func (v *FromSecret) UnmarshalYAML(unmarshal func(interface{}) error) error {
	err := unmarshal(&v.Name)
	if err == nil {
		return nil
	}

	s := struct {
		Name string
		Path string
	}{}
	err = unmarshal(&s)
	v.Name = s.Name
	v.Path = s.Path
	return err
}
