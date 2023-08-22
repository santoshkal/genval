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

	var dockerfile strings.Builder
	for _, stageData := range config.Dockerfile {
		dockerfile.WriteString(fmt.Sprintf("# Stage: %d\n", stageData.Stage))
		dockerfile.WriteString(fmt.Sprintf("FROM %s\n", stageData.From))

		fields := []struct {
			Condition bool
			Prefix    string
			Values    []string
		}{
			{stageData.Workdir != "", "WORKDIR", []string{stageData.Workdir}},
			{stageData.Arg != "", "ARG", []string{stageData.Arg}},
			{stageData.Env != "", "ENV", []string{stageData.Env}},
			{stageData.Label != "", "LABEL", []string{stageData.Label}},
			{stageData.Maintainer != "", "MAINTAINER", []string{stageData.Maintainer}},
			{stageData.User != "", "USER", []string{stageData.User}},
			{stageData.Volume != "", "VOLUME", []string{stageData.Volume}},
			{stageData.OnBuild != "", "ONBUILD", []string{stageData.OnBuild}},
			{stageData.StopSignal != "", "STOPSIGNAL", []string{stageData.StopSignal}},
			{stageData.Healthcheck != "", "HEALTHCHECK", []string{stageData.Healthcheck}},
		}

		for _, field := range fields {
			if field.Condition {
				dockerfile.WriteString(fmt.Sprintf("%s %s\n", field.Prefix, strings.Join(field.Values, " ")))
			}
		}

		for _, copy := range stageData.Copy {
			if copy != "" {
				dockerfile.WriteString(fmt.Sprintf("COPY %s\n", copy))
			}
		}

		for _, run := range stageData.Run {
			if run != "" {
				dockerfile.WriteString(fmt.Sprintf("RUN %s\n", run))
			}
		}

		if len(stageData.Entrypoint) > 0 {
			dockerfile.WriteString(fmt.Sprintf("ENTRYPOINT [\"%s\"]\n", strings.Join(stageData.Entrypoint, "\",\"")))
		}
		if len(stageData.Cmd) > 0 {
			dockerfile.WriteString(fmt.Sprintf("CMD [\"%s\"]\n", strings.Join(stageData.Cmd, "\",\"")))
		}
		if len(stageData.Add) > 0 {
			dockerfile.WriteString(fmt.Sprintf("ADD %s\n", strings.Join(stageData.Add, " ")))
		}
		if len(stageData.Shell) > 0 {
			dockerfile.WriteString(fmt.Sprintf("SHELL [\"%s\"]\n", strings.Join(stageData.Shell, "\",\"")))
		}

		dockerfile.WriteString("\n") // Add an empty line between stages
	}

	return dockerfile.String(), nil
}
