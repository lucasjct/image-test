package tests_test

import (
	"github.com/gruntwork-io/terratest/modules/docker"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Check node version from dockerfile", func() {
	It("should match the version", func() {

		tag := "undistrotest/docker_node_alpine_test"
		buildOptions := &docker.BuildOptions{
			Tags: []string{tag},
		}

		docker.Build(GinkgoT(), "../images/docker_node_alpine", buildOptions)

		opts := &docker.RunOptions{Command: []string{"node", "--version"}}
		output := docker.Run(GinkgoT(), tag, opts)
		//assert.Equal(GinkgoT(), "v20.18.1", output)
		Expect("v20.18.1").To(Equal(output))
	})

})
