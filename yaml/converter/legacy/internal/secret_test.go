// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package yaml

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestUnmarshalSecrets(t *testing.T) {
	testdata := []struct {
		from string
		want []*Secret
	}{
		{
			from: "[ mysql_username, mysql_password]",
			want: []*Secret{
				{
					Source: "mysql_username",
					Target: "mysql_username",
				},
				{
					Source: "mysql_password",
					Target: "mysql_password",
				},
			},
		},
		{
			from: "[ { source: mysql_prod_username, target: mysql_username } ]",
			want: []*Secret{
				{
					Source: "mysql_prod_username",
					Target: "mysql_username",
				},
			},
		},
		{
			from: "[ { source: mysql_prod_username, target: mysql_username }, { source: redis_username, target: redis_username } ]",
			want: []*Secret{
				{
					Source: "mysql_prod_username",
					Target: "mysql_username",
				},
				{
					Source: "redis_username",
					Target: "redis_username",
				},
			},
		},
	}

	for _, test := range testdata {
		in := []byte(test.from)
		got := Secrets{}
		err := yaml.Unmarshal(in, &got)
		if err != nil {
			t.Error(err)
		} else if !reflect.DeepEqual(test.want, got.Secrets) {
			t.Errorf("got secret %v want %v", got.Secrets, test.want)
		}
	}
}
