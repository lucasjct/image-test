package tests_test

import (
	"github.com/gruntwork-io/terratest/modules/docker"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Check network connection with Python image", func() {
	It("should return 200", func() {
		GinkgoT().Parallel()

		tag := "undistrotest/docker_python_test"
		buildOptions := &docker.BuildOptions{
			Tags: []string{tag},
		}

		docker.Build(GinkgoT(), "../images/docker_python", buildOptions)

		opts := &docker.RunOptions{Command: []string{"python", "send_request.py"}}
		output := docker.Run(GinkgoT(), tag, opts)
		// assert.Equal(GinkgoT(), "200", output)
		Expect("200").To(Equal(output))
	})

})