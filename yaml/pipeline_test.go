package yaml

import "testing"

func TestPipelineUnmarshal(t *testing.T) {
	diff, err := diff("testdata/pipeline.yml")
	if err != nil {
		t.Error(err)
	}
	if diff != "" {
		t.Error("Failed to parse pipeline")
		t.Log(diff)
	}
}
