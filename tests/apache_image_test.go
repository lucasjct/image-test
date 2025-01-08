package tests_test

import (
	"github.com/gruntwork-io/terratest/modules/docker"
	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Check apache variables and volume into the container", func() {

	var (
		tag          string
		buildOptions *docker.BuildOptions
	)
	BeforeEach(func() {
		// Define the image to test
		tag = "apache-env-test"

		// Build the Docker image
		buildOptions = &docker.BuildOptions{
			Tags:          []string{tag},
			Architectures: []string{"linux/amd64", "linux/arm64"},
		}
		docker.Build(GinkgoT(), "../images/docker_apache/", buildOptions)

	})
	Context("When running to check variables", func() {
		It("the container should return the correct variables", func() {
			GinkgoT().Parallel()

			// Define environment variables and their expected values
			envVars := map[string]string{
				"APACHE_LOCK_DIR": "/var/lock",
				"APACHE_PID_FILE": "/var/run/apache2.pid",
				"APACHE_RUN_USER": "www-data",
				"APACHE_RUN_GROU": "www-data",
				"APACHE_LOG_DIR=": "/var/log/apache2",
			}

			// Construct a shell command to print the environment variables
			command := "sh -c '"
			for key := range envVars {
				command += "echo $" + key + " && "
			}
			command = command[:len(command)-4] + "'" // Remove trailing " && "

			// Run the container with the constructed command
			runOptions := &docker.RunOptions{
				Command: []string{"sh", "-c", command},
			}
			output := docker.Run(GinkgoT(), tag, runOptions)

			// Validate each environment variable's value in the output
			for key, expectedValue := range envVars {
				assert.Contains(GinkgoT(), output, expectedValue, key, expectedValue)

			}
		})

		Context("When running the container to check the volume created", func() {
			It("the 'html' volume should be present on container ", func() {
				tag := "apache_volume_test"
				buildOptions := &docker.BuildOptions{
					Tags:          []string{tag},
					Architectures: []string{"linux/amd64", "linux/arm64"},
				}
				docker.Build(GinkgoT(), "../images/docker_apache/", buildOptions)
				opts := &docker.RunOptions{Command: []string{"ls", "/var/www/html"}}
				output := docker.Run(GinkgoT(), tag, opts)
				Expect("index.html").To(Equal(output))

			})
		})
	})
})
