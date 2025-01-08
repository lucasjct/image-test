package tests

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"gotest.tools/v3/assert"
)

func TestPythonImage(t *testing.T) {
	t.Parallel()

	tag := "undistrotest/docker_python_test"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	docker.Build(t, "../images/docker_python", buildOptions)

	opts := &docker.RunOptions{Command: []string{"python", "send_request.py"}}
	output := docker.Run(t, tag, opts)
	assert.Equal(t, "200", output)

}
