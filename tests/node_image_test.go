package tests_test

import (
	"github.com/gruntwork-io/terratest/modules/docker"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Check node version from dockerfile", func() {
	var (
		tag          string
		buildOptions *docker.BuildOptions
	)

	BeforeEach(func() {
		tag = "undistrotest/docker_node_alpine_test"
		buildOptions = &docker.BuildOptions{
			Tags: []string{tag},
		}

		docker.Build(GinkgoT(), "../images/docker_node_alpine", buildOptions)
	})

	When("running the container", func() {
		It("should match the version", func() {
			opts := &docker.RunOptions{Command: []string{"node", "--version"}}
			output := docker.Run(GinkgoT(), tag, opts)
			Expect("v20.18.1").To(Equal(output))
		})
	})
})
