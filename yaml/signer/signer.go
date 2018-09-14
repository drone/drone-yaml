package signer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/drone/drone-yaml/yaml"
)

// ErrInvalidKey is returned when the key is missing or
// is less than 32-bytes.
var ErrInvalidKey = errors.New("signer: key must be 32-bytes")

// Sign calculates and returns the hmac signature of the
// parsed yaml file.
func Sign(manifest *yaml.Manifest, key []byte) (string, error) {
	hmac, err := sign(manifest, key)
	return hex.EncodeToString(hmac), err
}

// SignUpdate calculates the hmac signature of the parsed
// yaml file and adds a signature resource. If a signature
// resource already exists, it is replaced.
func SignUpdate(manifest *yaml.Manifest, key []byte) error {
	hmac, err := sign(manifest, key)
	if err != nil {
		return err
	}
	for _, r := range manifest.Resources {
		if s, ok := r.(*yaml.Signature); ok {
			s.Hmac = hex.EncodeToString(hmac)
			return nil
		}
	}
	manifest.Resources = append(
		manifest.Resources,
		&yaml.Signature{
			Kind: "signature",
			Hmac: hex.EncodeToString(hmac),
		},
	)
	return nil
}

// Verify returns true if the signature of the parsed
// yaml file can be verified.
func Verify(manifest *yaml.Manifest, key []byte) (bool, error) {
	mac1, err := extract(manifest)
	if err != nil {
		return false, nil
	}
	mac2, err := sign(manifest, key)
	if err != nil {
		return false, err
	}
	return hmac.Equal(mac1, mac2), nil
}

// helper function extracts the hex-encoded signature
// resource from the parsed resource list.
func extract(manifest *yaml.Manifest) ([]byte, error) {
	for _, r := range manifest.Resources {
		if s, ok := r.(*yaml.Signature); ok {
			return hex.DecodeString(s.Hmac)
		}
	}
	return nil, errors.New("yaml: missing signature")
}

// helper function generates a hex-encoded signature
// based on the parsed resource list.
func sign(manifest *yaml.Manifest, key []byte) ([]byte, error) {
	if len(key) < 32 {
		return nil, ErrInvalidKey
	}

	var filtered []yaml.Resource
	for _, resource := range manifest.Resources {
		if _, ok := resource.(*yaml.Signature); !ok {
			filtered = append(filtered, resource)
		}
	}

	m := &yaml.Manifest{
		Resources: filtered,
	}
	h := hmac.New(sha256.New, key)
	e := json.NewEncoder(h).Encode(m)
	return h.Sum(nil), e
}
