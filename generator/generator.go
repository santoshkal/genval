// dockerfilegen/generator.go
package generator

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

// DockerfileConfig represents the structure of the YAML configuration.
type DockerfileConfig struct {
	Dockerfile []DockerfileData `yaml:"dockerfile"`
}

type DockerfileData struct {
	From            string   `yaml:"from"`
	Workdir         string   `yaml:"workdir"`
	Copy            []string `yaml:"copy"`
	Run             []string `yaml:"run"`
	Entrypoint      []string `yaml:"entrypoint"`
	Arg             string   `yaml:"arg"`
	Env             string   `yaml:"env"`
	Label           string   `yaml:"label"`
	Maintainer      string   `yaml:"maintainer"`
	Cmd             []string `yaml:"cmd"`
	Expose          int      `yaml:"expose"`
	User            string   `yaml:"user"`
	Add             []string `yaml:"add"`
	Volume          string   `yaml:"volume"`
	OnBuild         string   `yaml:"onbuild"`
	StopSignal      string   `yaml:"stopsignal"`
	Healthcheck     string   `yaml:"healthcheck"`
	Shell           []string `yaml:"shell"`
	Stage           int      `json:"_"`
	ProcessingStage int      `json:"stage"`
}

// GenerateDockerfile generates a Dockerfile content based on the provided YAML data.
func GenerateDockerfile(yamlData []byte) (string, error) {
	var config DockerfileConfig
	err := yaml.Unmarshal(yamlData, &config)
	if err != nil {
		return "", fmt.Errorf("error parsing YAML: %v", err)
	}

	dockerfile := fmt.Sprintf("FROM %s\n", config.Dockerfile[0].From)
	dockerfile += fmt.Sprintf("WORKDIR %s\n", config.Dockerfile[0].Workdir)
	if config.Dockerfile[0].Arg != "" {
		dockerfile += fmt.Sprintf("ARG %s\n", config.Dockerfile[0].Arg)
	}
	if config.Dockerfile[0].Env != "" {
		dockerfile += fmt.Sprintf("ENV %s\n", config.Dockerfile[0].Env)
	}
	if config.Dockerfile[0].Label != "" {
		dockerfile += fmt.Sprintf("LABEL %s\n", config.Dockerfile[0].Label)
	}
	if config.Dockerfile[0].Maintainer != "" {
		dockerfile += fmt.Sprintf("MAINTAINER %s\n", config.Dockerfile[0].Maintainer)
	}
	if config.Dockerfile[0].User != "" {
		dockerfile += fmt.Sprintf("USER %s\n", config.Dockerfile[0].User)
	}
	if config.Dockerfile[0].Volume != "" {
		dockerfile += fmt.Sprintf("VOLUME %s\n", config.Dockerfile[0].Volume)
	}
	if config.Dockerfile[0].OnBuild != "" {
		dockerfile += fmt.Sprintf("ONBUILD %s\n", config.Dockerfile[0].OnBuild)
	}
	if config.Dockerfile[0].StopSignal != "" {
		dockerfile += fmt.Sprintf("STOPSIGNAL %s\n", config.Dockerfile[0].StopSignal)
	}
	if config.Dockerfile[0].Healthcheck != "" {
		dockerfile += fmt.Sprintf("HEALTHCHECK %s\n", config.Dockerfile[0].Healthcheck)
	}

	for _, copy := range config.Dockerfile[0].Copy {
		if copy != "" {
			dockerfile += fmt.Sprintf("COPY %s\n", copy)
		}
	}

	for _, run := range config.Dockerfile[0].Run {
		if run != "" {
			dockerfile += fmt.Sprintf("RUN %s\n", run)
		}
	}

	if len(config.Dockerfile[0].Entrypoint) > 0 {
		dockerfile += fmt.Sprintf("ENTRYPOINT [\"%s\"]\n", strings.Join(config.Dockerfile[0].Entrypoint, "\",\""))
	}
	if len(config.Dockerfile[0].Cmd) > 0 {
		dockerfile += fmt.Sprintf("CMD [\"%s\"]\n", strings.Join(config.Dockerfile[0].Cmd, "\",\""))
	}
	if len(config.Dockerfile[0].Add) > 0 {
		dockerfile += fmt.Sprintf("ADD %s\n", config.Dockerfile[0].Add)
	}
	if len(config.Dockerfile[0].Shell) > 0 {
		dockerfile += fmt.Sprintf("SHELL %s\n", strings.Join(config.Dockerfile[0].Shell, "\",\""))

	}
	return dockerfile, nil
}
