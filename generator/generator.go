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
	From        string   `yaml:"from"`
	Workdir     string   `yaml:"workdir"`
	Copy        []string `yaml:"copy"`
	Run         []string `yaml:"run"`
	Entrypoint  []string `yaml:"entrypoint"`
	Arg         string   `yaml:"arg"`
	Env         string   `yaml:"env"`
	Label       string   `yaml:"label"`
	Maintainer  string   `yaml:"maintainer"`
	Cmd         []string `yaml:"cmd"`
	Expose      int      `yaml:"expose"`
	User        string   `yaml:"user"`
	Add         []string `yaml:"add"`
	Volume      string   `yaml:"volume"`
	OnBuild     string   `yaml:"onbuild"`
	StopSignal  string   `yaml:"stopsignal"`
	Healthcheck string   `yaml:"healthcheck"`
	Shell       []string `yaml:"shell"`
	Stage       int      `json:"_"`
}

// GenerateDockerfile generates a Dockerfile content based on the provided YAML data.
func GenerateDockerfile(yamlData []byte) (string, error) {
	var config DockerfileConfig

	err := yaml.Unmarshal(yamlData, &config)
	if err != nil {
		return "", fmt.Errorf("error parsing YAML: %v", err)
	}

	// Move the flags out as variables
	var dockerfile strings.Builder
	baseFlags := "as build"
	for _, stageData := range config.Dockerfile {
		dockerfile.WriteString(fmt.Sprintf("# Stage: %d\n", stageData.Stage))
		if stageData.Stage == 0 && stageData.From != "" {
			dockerfile.WriteString(fmt.Sprintf("FROM %s %s\n", stageData.From, baseFlags))
		} else {

			dockerfile.WriteString(fmt.Sprintf("FROM %s\n", stageData.From))
		}

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

		copyFlags := "--chown=65532:65532 --from=build "

		for _, copy := range stageData.Copy {
			if copy != "" && stageData.Stage > 0 {
				dockerfile.WriteString(fmt.Sprintf("COPY %s %s\n", copyFlags, copy))
			} else {
				dockerfile.WriteString(fmt.Sprintf("COPY %s\n", copy))
			}

		}

		var runInstructions []string
		if stageData.Stage == 0 && len(stageData.Run) > 1 {
			for _, run := range config.Dockerfile[0].Run {
				if run != "" {
					runInstructions = append(runInstructions, run)
				}
			}

			if len(runInstructions) > 0 {
				dockerfile.WriteString("RUN ")
				dockerfile.WriteString(strings.Join(runInstructions, " \\\n    && "))
				dockerfile.WriteString("\n")
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
