// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package pretty

import (
	"sort"
	"strings"

	"github.com/drone/drone-yaml/yaml"
)

// TODO consider "!!binary |" for secret value

// helper function to pretty prints the signature resource.
func printSecret(w writer, v *yaml.Secret) {
	w.WriteString("---")
	w.WriteTagValue("version", v.Version)
	w.WriteTagValue("kind", v.Kind)

	if len(v.Data) > 0 {
		w.WriteTagValue("type", toSecretType(v.Type))
		w.WriteTagValue("name", v.Name)
		printData(w, v.Data)
	}
	if len(v.External) > 0 {
		w.WriteTagValue("type", toSecretType(v.Type))
		w.WriteTagValue("name", v.Name)
		printExternalData(w, v.External)
	}
	if isSecretGetEmpty(v.Get) == false {
		w.WriteTagValue("type", v.Type)
		w.WriteTagValue("name", v.Name)
		w.WriteByte('\n')
		printGet(w, v.Get)
	}
	w.WriteByte('\n')
	w.WriteByte('\n')
}

// helper function returns the secret type text.
func toSecretType(s string) string {
	s = strings.ToLower(s)
	switch s {
	case "docker", "ecr", "general":
		return s
	default:
		return "general"
	}
}

// helper function prints the get block.
func printGet(w writer, v yaml.SecretGet) {
	w.WriteTag("get")
	w.IndentIncrease()
	w.WriteTagValue("path", v.Path)
	w.WriteTagValue("name", v.Name)
	w.WriteTagValue("key", v.Key)
	w.IndentDecrease()
}

// helper function prints the external data.
func printExternalData(w writer, d map[string]yaml.ExternalData) {
	var keys []string
	for k := range d {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	w.WriteTag("external_data")
	w.IndentIncrease()
	for _, k := range keys {
		v := d[k]
		w.WriteTag(k)
		w.IndentIncrease()
		w.WriteTagValue("path", v.Path)
		w.WriteTagValue("name", v.Name)
		w.IndentDecrease()
	}
	w.IndentDecrease()
}

func printData(w writer, d map[string]string) {
	var keys []string
	for k := range d {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	w.WriteTag("data")
	w.IndentIncrease()
	for _, k := range keys {
		v := d[k]
		w.WriteTag(k)
		w.WriteByte(' ')
		w.WriteByte('>')
		w.IndentIncrease()
		v = spaceReplacer.Replace(v)
		for _, s := range chunk(v, 60) {
			w.WriteByte('\n')
			w.Indent()
			w.WriteString(s)
		}
		w.IndentDecrease()
	}
	w.IndentDecrease()
}

// replace spaces and newlines.
var spaceReplacer = strings.NewReplacer(" ", "", "\n", "")

// helper function returns true if the secret get
// object is empty.
func isSecretGetEmpty(v yaml.SecretGet) bool {
	return v.Key == "" &&
		v.Name == "" &&
		v.Path == ""
}
