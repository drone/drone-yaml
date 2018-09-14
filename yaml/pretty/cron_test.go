package pretty

import "testing"

func TestCron(t *testing.T) {
	ok, err := diff("testdata/cron.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}
