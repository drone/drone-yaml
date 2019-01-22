package circleci

type (
	// Config defines the pipeline configuration.
	Config struct {
		// Version specifies the yaml configuration
		// file version.
		Version string

		// Jobs defines a list of pipeline jobs.
		Jobs []*Job

		// Workflows are used to orchestrate jobs.
		Workflows struct {
			Version string
			List    map[string]*Workflow `yaml:",inline"`
		}
	}

	// Workflow ochestrates one or more jobs.
	Workflow struct {
		Jobs []string
	}

	// Job defines a pipeline job.
	Job struct {
		// Name of the stage.
		Name string

		// Docker configures a Docker executor.
		Docker Docker

		// Environment variables passed to the executor.
		Environment map[string]string

		// Steps configures the Job steps.
		Steps map[string]Step

		// Branches limits execution by branch.
		Branches []struct {
			Only   []string
			Ignore []string
		}
	}

	// Step defines a build execution unit.
	Step struct {
		Run                Run
		AddSSHKeys         map[string]interface{} `yaml:"add_ssh_keys"`
		AttachWorkspace    map[string]interface{} `yaml:"attach_workspace"`
		Checkout           map[string]interface{} `yaml:"checkout"`
		Deploy             map[string]interface{} `yaml:"deploy"`
		PersistToWorkspace map[string]interface{} `yaml:"persist_to_workspace"`
		RestoreCache       map[string]interface{} `yaml:"restore_cache"`
		SaveCache          map[string]interface{} `yaml:"save_cache"`
		SetupRemoteDocker  map[string]interface{} `yaml:"setup_remote_docker"`
		StoreArtifacts     map[string]interface{} `yaml:"store_artifacts"`
		StoreTestResults   map[string]interface{} `yaml:"store_test_results"`
	}
)

// // UnmarshalYAML implements custom parsing for the stage section of the yaml
// // to cleanup the structure a bit.
// func (s *Stage) UnmarshalYAML(unmarshal func(interface{}) error) error {
// 	in := []struct {
// 		Step *Step
// 	}{}
// 	err := unmarshal(&in)
// 	if err != nil {
// 		return err
// 	}
// 	for _, step := range in {
// 		s.Steps = append(s.Steps, step.Step)
// 	}
// 	return nil
// }
