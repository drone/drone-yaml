package yaml

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestPort(t *testing.T) {
	tests := []struct {
		yaml     string
		port     int
		host     int
		protocol string
	}{
		{
			yaml: "80",
			port: 80,
		},
		{
			yaml:     "{ port: 80, host: 8080, protocol: TCP }",
			port:     80,
			host:     8080,
			protocol: "TCP",
		},
	}
	for _, test := range tests {
		in := []byte(test.yaml)
		out := new(Port)
		err := yaml.Unmarshal(in, out)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := out.Port, test.port; got != want {
			t.Errorf("Want Port %d, got %d", want, got)
		}
		if got, want := out.Host, test.host; got != want {
			t.Errorf("Want Host %d, got %d", want, got)
		}
		if got, want := out.Protocol, test.protocol; got != want {
			t.Errorf("Want Host %s, got %s", want, got)
		}
	}
}
