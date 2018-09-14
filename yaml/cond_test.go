package yaml

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestConstraintMatch(t *testing.T) {
	testdata := []struct {
		conf string
		with string
		want bool
	}{
		// slice value
		{
			conf: "[ master, feature/* ]",
			with: "develop",
			want: false,
		},
		{
			conf: "[ master, feature/* ]",
			with: "master",
			want: true,
		},
		{
			conf: "[ master, feature/* ]",
			with: "feature/foo",
			want: true,
		},
		// includes block
		{
			conf: "include: [ master ]",
			with: "develop",
			want: false,
		},
		{
			conf: "include: [ master] ",
			with: "master",
			want: true,
		},
		{
			conf: "include: [ feature/* ]",
			with: "master",
			want: false,
		},
		{
			conf: "include: [ feature/* ]",
			with: "feature/foo",
			want: true,
		},
		{
			conf: "include: [ master, feature/* ]",
			with: "develop",
			want: false,
		},
		{
			conf: "include: [ master, feature/* ]",
			with: "master",
			want: true,
		},
		{
			conf: "include: [ master, feature/* ]",
			with: "feature/foo",
			want: true,
		},
		// excludes block
		{
			conf: "exclude: [ master ]",
			with: "develop",
			want: true,
		},
		{
			conf: "exclude: [ master ]",
			with: "master",
			want: false,
		},
		{
			conf: "exclude: [ feature/* ]",
			with: "master",
			want: true,
		},
		{
			conf: "exclude: [ feature/* ]",
			with: "feature/foo",
			want: false,
		},
		{
			conf: "exclude: [ master, develop ]",
			with: "master",
			want: false,
		},
		{
			conf: "exclude: [ feature/*, bar ]",
			with: "master",
			want: true,
		},
		{
			conf: "exclude: [ feature/*, bar ]",
			with: "feature/foo",
			want: false,
		},
		// include and exclude blocks
		{
			conf: "{ include: [ master, feature/* ], exclude: [ develop ] }",
			with: "master",
			want: true,
		},
		{
			conf: "{ include: [ master, feature/* ], exclude: [ feature/bar ] }",
			with: "feature/bar",
			want: false,
		},
		{
			conf: "{ include: [ master, feature/* ], exclude: [ master, develop ] }",
			with: "master",
			want: false,
		},
		// empty blocks
		{
			conf: "",
			with: "master",
			want: true,
		},
		// double star
		{
			conf: "foo/**",
			with: "foo/bar/baz/qux",
			want: true,
		},
		{
			conf: "foo/**/qux",
			with: "foo/bar/baz/qux",
			want: true,
		},
	}
	for _, test := range testdata {
		c := parseCondition(test.conf)
		got, want := c.Match(test.with), test.want
		if got != want {
			t.Errorf("Expect %q matches %q is %v", test.with, test.conf, want)
		}
	}
}

func parseCondition(s string) *Condition {
	c := &Condition{}
	yaml.Unmarshal([]byte(s), c)
	return c
}
