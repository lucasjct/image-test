package tests

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"gotest.tools/v3/assert"
)

func TestDockerAlpineNode(t *testing.T) {

	tag := "undistrotest/docker_node_alpine_test"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	docker.Build(t, "../images/docker_node_alpine", buildOptions)

	opts := &docker.RunOptions{Command: []string{"node", "--version"}}
	output := docker.Run(t, tag, opts)
	assert.Equal(t, "v20.18.1", output)

}
