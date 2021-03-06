// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package yaml

type (
	// Secrets represents a list of container secrets.
	Secrets struct {
		Secrets []*Secret
	}

	// Secret represents a container secret.
	Secret struct {
		Source string
		Target string
	}
)

// UnmarshalYAML implements the Unmarshaller interface.
func (s *Secrets) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var strslice []string
	err := unmarshal(&strslice)
	if err == nil {
		for _, str := range strslice {
			s.Secrets = append(s.Secrets, &Secret{
				Source: str,
				Target: str,
			})
		}
		return nil
	}
	return unmarshal(&s.Secrets)
}
