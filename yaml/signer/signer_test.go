package signer

import (
	"testing"

	"github.com/drone/drone-yaml/yaml"
)

func TestSign(t *testing.T) {
	manifest, err := yaml.ParseFile("testdata/signed.yml")
	if err != nil {
		t.Error(err)
		return
	}
	key := []byte("589396227fff5a93ba934965e8735f88")
	want := "43ea63152d72a554b2ab2bba0bac0d33d5d6d2de6f368c0c6ff544981f1f94ef"
	got, err := Sign(manifest, key)
	if err != nil {
		t.Error(err)
	}
	if got != want {
		t.Errorf("Expected hash %s, got %s", want, got)
	}

	verified, err := Verify(manifest, key)
	if err != nil {
		t.Error(err)
	}
	if !verified {
		t.Errorf("Expected signature verified")
	}
}

func TestVerify_Failure(t *testing.T) {
	manifest, err := yaml.ParseFile("testdata/signed.yml")
	if err != nil {
		t.Error(err)
		return
	}
	key := []byte("c953bd41ad0f75848a78ccd54d3861fa")
	verified, err := Verify(manifest, key)
	if err != nil {
		t.Error(err)
	}
	if verified {
		t.Errorf("Expected signature verification failure")
	}
}

func TestVerify_InvalidKey(t *testing.T) {
	manifest, err := yaml.ParseFile("testdata/signed.yml")
	if err != nil {
		t.Error(err)
		return
	}
	key := []byte("this-is-an-invalid-key")
	_, err = Verify(manifest, key)
	if err != ErrInvalidKey {
		t.Errorf("Expected ErrInvalidKey")
	}
}

// This test verifies that signature verification
// fails if no signature is present in the yaml.
func TestVerify_MissingSignature(t *testing.T) {
	manifest, err := yaml.ParseFile("testdata/missing_signature.yml")
	if err != nil {
		t.Error(err)
		return
	}
	key := []byte("589396227fff5a93ba934965e8735f88")
	verified, err := Verify(manifest, key)
	if err != nil {
		t.Error(err)
	}
	if verified {
		t.Errorf("Expected signature verification failure")
	}
}

// The test verifies that the SignUpdate function signs the
// configuraiton and appends the signature.
func TestSignUpdate(t *testing.T) {
	manifest, err := yaml.ParseFile("testdata/signed.yml")
	if err != nil {
		t.Error(err)
		return
	}
	// verify the yaml defines 3 resources. 1 pipeline,
	// 1 secret, and 1 signature.
	if len(manifest.Resources) != 3 {
		t.Errorf("Expect 3 resources in yaml file")
	}

	key := []byte("589396227fff5a93ba934965e8735f88")

	SignUpdate(manifest, key)
	signature, ok := manifest.Resources[len(manifest.Resources)-1].(*yaml.Signature)
	if !ok {
		t.Errorf("Expected signature appended to the resource list")
	}

	// verify that the existing signature resource
	// has been replaced.
	if len(manifest.Resources) != 3 {
		t.Errorf("Expect 3 resources in yaml file")
	}

	hash := "43ea63152d72a554b2ab2bba0bac0d33d5d6d2de6f368c0c6ff544981f1f94ef"
	if got, want := signature.Hmac, hash; got != want {
		t.Errorf("Expected hash %s, got %s", want, got)
	}

	verified, err := Verify(manifest, key)
	if err != nil {
		t.Error(err)
	}
	if !verified {
		t.Errorf("Expected signature verified")
	}
}

// The test verifies that the SignUpdate function signs the
// configuraiton and appends the signature.
func TestSignUpdate_Append(t *testing.T) {
	manifest, err := yaml.ParseFile("testdata/signed.yml")
	if err != nil {
		t.Error(err)
		return
	}

	// verify the yaml defines 3 resources. 1 pipeline,
	// 1 secret, and 1 signature.
	if len(manifest.Resources) != 3 {
		t.Errorf("Expect 3 resources in yaml file")
	}

	// remove the last resource from the list, which
	// should be the signature resource.
	manifest.Resources = manifest.Resources[:len(manifest.Resources)-1]
	if len(manifest.Resources) != 2 {
		t.Errorf("Expect 2 resources in yaml file")
	}

	key := []byte("589396227fff5a93ba934965e8735f88")

	SignUpdate(manifest, key)
	signature, ok := manifest.Resources[len(manifest.Resources)-1].(*yaml.Signature)
	if !ok {
		t.Errorf("Expected signature appended to the resource list")
	}

	// verify that the signature resource has been
	// appended to the resource list.
	if len(manifest.Resources) != 3 {
		t.Errorf("Expect 3 resources in yaml file")
	}

	hash := "43ea63152d72a554b2ab2bba0bac0d33d5d6d2de6f368c0c6ff544981f1f94ef"
	if got, want := signature.Hmac, hash; got != want {
		t.Errorf("Expected hash %s, got %s", want, got)
	}

	verified, err := Verify(manifest, key)
	if err != nil {
		t.Error(err)
	}
	if !verified {
		t.Errorf("Expected signature verified")
	}
}

func TestSignUpdate_InvalidKey(t *testing.T) {
	manifest, err := yaml.ParseFile("testdata/signed.yml")
	if err != nil {
		t.Error(err)
		return
	}
	key := []byte("this-is-an-invalid-key")
	err = SignUpdate(manifest, key)
	if err != ErrInvalidKey {
		t.Errorf("Expected ErrInvalidKey")
	}
}
