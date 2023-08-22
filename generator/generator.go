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

	dockerfile := ""
	for _, stageData := range config.Dockerfile {
		dockerfile += fmt.Sprintf("# Stage: %d\n", stageData.Stage)
		dockerfile += fmt.Sprintf("FROM %s\n", stageData.From)
		if stageData.Workdir != "" {
			dockerfile += fmt.Sprintf("WORKDIR %s\n", stageData.Workdir)
		}
		if stageData.Arg != "" {
			dockerfile += fmt.Sprintf("ARG %s\n", stageData.Arg)
		}
		if stageData.Env != "" {
			dockerfile += fmt.Sprintf("ENV %s\n", stageData.Env)
		}
		if stageData.Label != "" {
			dockerfile += fmt.Sprintf("LABEL %s\n", stageData.Label)
		}
		if stageData.Maintainer != "" {
			dockerfile += fmt.Sprintf("MAINTAINER %s\n", stageData.Maintainer)
		}
		if stageData.User != "" {
			dockerfile += fmt.Sprintf("USER %s\n", stageData.User)
		}
		if stageData.Volume != "" {
			dockerfile += fmt.Sprintf("VOLUME %s\n", stageData.Volume)
		}
		if stageData.OnBuild != "" {
			dockerfile += fmt.Sprintf("ONBUILD %s\n", stageData.OnBuild)
		}
		if stageData.StopSignal != "" {
			dockerfile += fmt.Sprintf("STOPSIGNAL %s\n", stageData.StopSignal)
		}
		if stageData.Healthcheck != "" {
			dockerfile += fmt.Sprintf("HEALTHCHECK %s\n", stageData.Healthcheck)
		}

		for _, copy := range stageData.Copy {
			if copy != "" {
				dockerfile += fmt.Sprintf("COPY %s\n", copy)
			}
		}

		for _, run := range stageData.Run {
			if run != "" {
				dockerfile += fmt.Sprintf("RUN %s\n", run)
			}
		}

		if len(stageData.Entrypoint) > 0 {
			dockerfile += fmt.Sprintf("ENTRYPOINT [\"%s\"]\n", strings.Join(stageData.Entrypoint, "\",\""))
		}
		if len(stageData.Cmd) > 0 {
			dockerfile += fmt.Sprintf("CMD [\"%s\"]\n", strings.Join(stageData.Cmd, "\",\""))
		}
		if len(stageData.Add) > 0 {
			dockerfile += fmt.Sprintf("ADD %s\n", strings.Join(stageData.Add, " "))
		}
		if len(stageData.Shell) > 0 {
			dockerfile += fmt.Sprintf("SHELL [\"%s\"]\n", strings.Join(stageData.Shell, "\",\""))
		}
		dockerfile += "\n" // Add an empty line between stages
	}

	return dockerfile, nil
}
