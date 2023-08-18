// dockerfilegen/generator.go
package generator

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

// DockerfileConfig represents the structure of the YAML configuration.
type DockerfileConfig struct {
	Dockerfile DockerfileData `yaml:"dockerfile"`
}

type DockerfileData struct {
	From       string   `yaml:"from"`
	Workdir    string   `yaml:"workdir"`
	Copy       []string `yaml:"copy"`
	Run        []string `yaml:"run"`
	Entrypoint []string `yaml:"entrypoint"`
}

// GenerateDockerfile generates a Dockerfile content based on the provided YAML data.
func GenerateDockerfile(yamlData []byte) (string, error) {
	var config DockerfileConfig
	err := yaml.Unmarshal(yamlData, &config)
	if err != nil {
		return "", fmt.Errorf("error parsing YAML: %v", err)
	}

	dockerfile := fmt.Sprintf("FROM %s\n", config.Dockerfile.From)
	dockerfile += fmt.Sprintf("WORKDIR %s\n", config.Dockerfile.Workdir)

	for _, copy := range config.Dockerfile.Copy {
		dockerfile += fmt.Sprintf("COPY %s\n", copy)
	}

	for _, run := range config.Dockerfile.Run {
		dockerfile += fmt.Sprintf("RUN %s\n", run)
	}

	if len(config.Dockerfile.Entrypoint) > 0 {
		dockerfile += fmt.Sprintf("ENTRYPOINT [\"%s\"]\n", strings.Join(config.Dockerfile.Entrypoint, "\",\""))
	}

	// Add more logic for other Dockerfile instructions

	return dockerfile, nil
}
