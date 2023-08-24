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

// type Copy struct {
// 	Src  string `yaml:"src"`
// 	Dest string `yaml:"dest"`
// }

type DockerfileData struct {
	From        string   `yaml:"from"`
	Env         string   `yaml:"env"`
	Env1        string   `yaml:"env1"`
	Env2        string   `yaml:"env2"`
	Workdir     string   `yaml:"workdir"`
	Copy        []string `yaml:"copy"`
	CopyCmd     []string `yaml:"copyCmd"`
	Run         []string `yaml:"run"`
	RunCmd      []string `yaml:"runCmd"`
	Entrypoint  []string `yaml:"entrypoint"`
	Arg         string   `yaml:"arg"`
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

		copyFlags := "--chown=65532:65532 --from=build "

		if len(stageData.Cmd) > 0 {
			dockerfile.WriteString(fmt.Sprintf("CMD [\"%s\"]\n", strings.Join(stageData.Cmd, "\",\"")))
		}
		if len(stageData.Add) > 0 {
			dockerfile.WriteString(fmt.Sprintf("ADD %s\n", strings.Join(stageData.Add, " ")))
		}
		if len(stageData.Shell) > 0 {
			dockerfile.WriteString(fmt.Sprintf("SHELL [\"%s\"]\n", strings.Join(stageData.Shell, "\",\"")))
		}

		fieldsOrder := getFieldOrder(stageData)

		for _, fieldName := range fieldsOrder {
			switch fieldName {
			case "Workdir":
				dockerfile.WriteString(fmt.Sprintf("WORKDIR %s\n", stageData.Workdir))
			case "Arg":
				dockerfile.WriteString(fmt.Sprintf("ARG %s\n", stageData.Arg))
			case "Env":
				dockerfile.WriteString(fmt.Sprintf("ENV %s\n", stageData.Env))
			case "Copy":
				for _, copy := range stageData.Copy {
					if copy != "" && stageData.Stage > 0 {
						dockerfile.WriteString(fmt.Sprintf("COPY %s %s\n", copyFlags, copy))
					} else {
						dockerfile.WriteString(fmt.Sprintf("COPY %s\n", copy))
					}

				}
			case "CopyCmd":
				for _, copyCmd := range stageData.CopyCmd {
					if copyCmd != "" && stageData.Stage > 0 {
						dockerfile.WriteString(fmt.Sprintf("COPY %s %s\n", copyFlags, copyCmd))
					} else {
						dockerfile.WriteString(fmt.Sprintf("COPY %s\n", copyCmd))
					}

				}
			case "Run":
				runInstructions := make([]string, 0)
				for _, runInstruction := range stageData.Run {
					if runInstruction != "" {
						runInstructions = append(runInstructions, runInstruction)
					}
				}
				if len(runInstructions) > 0 {
					dockerfile.WriteString("RUN ")
					dockerfile.WriteString(strings.Join(runInstructions, " \\\n    && "))
					dockerfile.WriteString("\n")
				}
			case "RunCmd":
				runCmdInstructions := make([]string, 0)
				for _, runCmdInstruction := range stageData.RunCmd {
					if runCmdInstruction != "" {
						runCmdInstructions = append(runCmdInstructions, runCmdInstruction)
					}
				}
				if len(runCmdInstructions) > 0 {
					dockerfile.WriteString("RUN ")
					dockerfile.WriteString(strings.Join(runCmdInstructions, " \\\n    && "))
					dockerfile.WriteString("\n")
				}
			case "Env1":
				dockerfile.WriteString(fmt.Sprintf("ENV %s\n", stageData.Env1))
			case "Env2":
				dockerfile.WriteString(fmt.Sprintf("ENV %s\n", stageData.Env2))
			case "Entrypoint":
				dockerfile.WriteString(fmt.Sprintf("ENTRYPOINT %s\n", strings.Join(stageData.Entrypoint, " ")))
			case "Label":
				dockerfile.WriteString(fmt.Sprintf("LABEL %s\n", stageData.Label))
			case "Maintainer":
				dockerfile.WriteString(fmt.Sprintf("MAINTAINER %s\n", stageData.Maintainer))
			case "Cmd":
				dockerfile.WriteString(fmt.Sprintf("CMD %s\n", strings.Join(stageData.Cmd, " ")))
			case "Expose":
				dockerfile.WriteString(fmt.Sprintf("EXPOSE %d\n", stageData.Expose))
			case "User":
				dockerfile.WriteString(fmt.Sprintf("USER %s\n", stageData.User))
			case "Add":
				dockerfile.WriteString(fmt.Sprintf("ADD %s\n", strings.Join(stageData.Add, " ")))
			case "Volume":
				dockerfile.WriteString(fmt.Sprintf("VOLUME %s\n", stageData.Volume))
			case "OnBuild":
				dockerfile.WriteString(fmt.Sprintf("ONBUILD %s\n", stageData.OnBuild))
			case "StopSignal":
				dockerfile.WriteString(fmt.Sprintf("STOPSIGNAL %s\n", stageData.StopSignal))
			case "Healthcheck":
				dockerfile.WriteString(fmt.Sprintf("HEALTHCHECK %s\n", stageData.Healthcheck))
			case "Shell":
				dockerfile.WriteString(fmt.Sprintf("SHELL %s\n", strings.Join(stageData.Shell, " ")))
			}
		}

		dockerfile.WriteString("\n") // Add an empty line between stages
	}

	return dockerfile.String(), nil
}
