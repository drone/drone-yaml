package circleci

// Docker configures a Docker executor.
type Docker struct {
	// Image is the Docker image name.
	Image string

	// Name is the Docker container hostname.
	Name string

	// Entrypoint is the Docker container entrypoint.
	Entrypoint []string

	// Command is the Docker container command.
	Command []string

	// User is user that runs the Docker entrypoint.
	User string

	// Environment variables passed to the container.
	Environment map[string]string

	// Auth credentials to pull private images.
	Auth map[string]string

	// Auth credentials to pull private ECR images.
	AWSAuth map[string]string `yaml:"aws_auth"`
}
