package pretty

import "testing"

// func TestPipeline(t *testing.T) {
// 	ok, err := diff("testdata/pipeline.yml")
// 	if err != nil {
// 		t.Error(err)
// 	} else if !ok {
// 		t.Errorf("Unepxected formatting")
// 	}
// }

func TestPipeline_Build_Short(t *testing.T) {
	ok, err := diff("testdata/pipeline_build_short.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}

func TestPipeline_Build_Long(t *testing.T) {
	ok, err := diff("testdata/pipeline_build_long.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}

func TestPipeline_Clone_Depth(t *testing.T) {
	ok, err := diff("testdata/pipeline_clone_depth.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}

func TestPipeline_Clone_Disable(t *testing.T) {
	ok, err := diff("testdata/pipeline_clone_disable.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}

func TestPipeline_Depends(t *testing.T) {
	ok, err := diff("testdata/pipeline_depends.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}

func TestPipeline_Push(t *testing.T) {
	ok, err := diff("testdata/pipeline_push.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}

func TestPipeline_Ports(t *testing.T) {
	ok, err := diff("testdata/pipeline_ports.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}

func TestPipeline_Resources(t *testing.T) {
	ok, err := diff("testdata/pipeline_resources.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}

func TestPipeline_Services(t *testing.T) {
	ok, err := diff("testdata/services.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}

func TestPipeline_Settings(t *testing.T) {
	ok, err := diff("testdata/pipeline_settings.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}

func TestPipeline_Trigger(t *testing.T) {
	ok, err := diff("testdata/pipeline_trigger.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}

func TestPipeline_Volumes(t *testing.T) {
	ok, err := diff("testdata/pipeline_volumes.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}

func TestPipeline_Workspace(t *testing.T) {
	ok, err := diff("testdata/pipeline_workspace.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}
