package signer

import (
	"io/ioutil"
	"testing"
)

func TestSign(t *testing.T) {
	in, err := ioutil.ReadFile("testdata/signed.yml")
	if err != nil {
		t.Error(err)
		return
	}

	key := KeyString("589396227fff5a93ba934965e8735f88")
	want := "389cf92a7472870783a9a5ea77b4abe58a4bb67ba58e1e7293e943aee314aedc"
	got, err := Sign(in, key)
	if err != nil {
		t.Error(err)
	}
	if got != want {
		t.Errorf("Expected hash %s, got %s", want, got)
	}

	verified, err := Verify(in, key)
	if err != nil {
		t.Error(err)
	}
	if !verified {
		t.Errorf("Expected signature verified")
	}
}

func TestVerify_Invalid(t *testing.T) {
	in, err := ioutil.ReadFile("testdata/invalid_signature.yml")
	if err != nil {
		t.Error(err)
		return
	}
	key := KeyString("c953bd41ad0f75848a78ccd54d3861fa")
	verified, err := Verify(in, key)
	if err != nil {
		t.Error(err)
	}
	if verified {
		t.Errorf("Expected signature verification failure")
	}
}

func TestVerify_InvalidKey(t *testing.T) {
	in, err := ioutil.ReadFile("testdata/signed.yml")
	if err != nil {
		t.Error(err)
		return
	}
	key := []byte("this-is-an-invalid-key")
	_, err = Verify(in, key)
	if err != ErrInvalidKey {
		t.Errorf("Expected ErrInvalidKey")
	}
}

// This test verifies that signature verification
// fails if no signature is present in the yaml.
func TestVerify_MissingSignature(t *testing.T) {
	in, err := ioutil.ReadFile("testdata/missing_signature.yml")
	if err != nil {
		t.Error(err)
		return
	}
	key := KeyString("589396227fff5a93ba934965e8735f88")
	verified, err := Verify(in, key)
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
	before, err := ioutil.ReadFile("testdata/invalid_signature.yml")
	if err != nil {
		t.Error(err)
		return
	}

	key := KeyString("589396227fff5a93ba934965e8735f88")
	after, err := SignUpdate(before, key)
	if err != nil {
		t.Error(err)
		return
	}

	verified, err := Verify(after, key)
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
	before, err := ioutil.ReadFile("testdata/missing_signature.yml")
	if err != nil {
		t.Error(err)
		return
	}

	key := KeyString("589396227fff5a93ba934965e8735f88")
	after, err := SignUpdate(before, key)
	if err != nil {
		t.Error(err)
		return
	}

	verified, err := Verify(after, key)
	if err != nil {
		t.Error(err)
	}
	if !verified {
		t.Errorf("Expected signature verified")
	}
}

func TestSignUpdate_InvalidKey(t *testing.T) {
	in, err := ioutil.ReadFile("testdata/signed.yml")
	if err != nil {
		t.Error(err)
		return
	}
	key := KeyString("this-is-an-invalid-key")
	_, err = SignUpdate(in, key)
	if err != ErrInvalidKey {
		t.Errorf("Expected ErrInvalidKey")
	}
}
