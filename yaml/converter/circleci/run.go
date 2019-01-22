package circleci

import "time"

// Run defines a command
type Run struct {
	// Name of the command
	Name string

	// Command run in the shell.
	Command string

	// Shell to use to execute the command.
	Shell string

	// Workiring Directory in which the command
	// is run.
	WorkingDir string `yaml:"working_directory"`

	// Command is run in the background.
	Background bool `yaml:"background"`

	// Amount of time the command can run with
	// no output before being canceled.
	NoOutputTimeout time.Duration `yaml:"no_output_timeout"`

	// Environment variables set when running
	// the command in the shell.
	Environment map[string]string

	// Defines when the command should be executed.
	// Values are always, on_success, and on_fail.
	When string
}
