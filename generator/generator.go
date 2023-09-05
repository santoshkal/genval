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
	Env         string   `yaml:"env"`
	Env1        string   `yaml:"env1"`
	Env2        string   `yaml:"env2"`
	Arg         string   `yaml:"arg"`
	Workdir     string   `yaml:"workdir"`
	Copy        []string `yaml:"copy"`
	Run         []string `yaml:"run"`
	CopyCmd     []string `yaml:"copyCmd"`
	RunCmd      []string `yaml:"runCmd"`
	CopyCmd1    []string `yaml:"copyCmd1"`
	RunCmd1     []string `yaml:"runCmd1"`
	Label       string   `yaml:"label"`
	Maintainer  string   `yaml:"maintainer"`
	Expose      int      `yaml:"expose"`
	User        string   `yaml:"user"`
	Add         []string `yaml:"add"`
	AddCmd      []string `yaml:"addcmd"`
	Volume      string   `yaml:"volume"`
	OnBuild     string   `yaml:"onbuild"`
	StopSignal  string   `yaml:"stopsignal"`
	Healthcheck string   `yaml:"healthcheck"`
	Entrypoint  []string `yaml:"entrypoint"`
	Cmd         []string `yaml:"cmd"`
	Shell       []string `yaml:"shell"`
	Stage       int      `json:"_"`
}

// func parseRunCommands(commands []string, dockerfile *strings.Builder) {
// 	if len(commands) > 0 {
// 		dockerfile.WriteString("RUN ")
// 		dockerfile.WriteString(strings.Join(commands, " \\\n    && "))
// 		dockerfile.WriteString("\n")
// 	}
// }

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
			{stageData.Arg != "", "ARG", []string{stageData.Arg}},
			{stageData.Env != "", "ENV", []string{stageData.Env}},
			{stageData.Env1 != "", "ENV", []string{stageData.Env1}},
			{stageData.Env2 != "", "ENV", []string{stageData.Env2}},
			{stageData.Label != "", "LABEL", []string{stageData.Label}},
			{stageData.Maintainer != "", "MAINTAINER", []string{stageData.Maintainer}},
			{stageData.Workdir != "", "WORKDIR", []string{stageData.Workdir}},
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

		if len(stageData.Add) > 0 {
			dockerfile.WriteString(fmt.Sprintf("ADD %s\n", strings.Join(stageData.Add, " ")))
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
		for _, copyCmd := range stageData.CopyCmd {
			if copyCmd != "" {
				dockerfile.WriteString(fmt.Sprintf("COPY %s\n", copyCmd))
			}

		}
		// parseRunCommands(stageData.Run, &dockerfile)
		// parseRunCommands(stageData.RunCmd, &dockerfile)
		// parseRunCommands(stageData.RunCmd1, &dockerfile)
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
		runCmd1Instructions := make([]string, 0)
		for _, runCmd1Instruction := range stageData.RunCmd1 {
			if runCmd1Instruction != "" {
				runCmd1Instructions = append(runCmd1Instructions, runCmd1Instruction)
			}
		}
		if len(runCmd1Instructions) > 0 {
			dockerfile.WriteString("RUN ")
			dockerfile.WriteString(strings.Join(runCmd1Instructions, " \\\n    && "))
			dockerfile.WriteString("\n")
		}

		if len(stageData.Shell) > 0 {
			dockerfile.WriteString(fmt.Sprintf("SHELL [\"%s\"]\n", strings.Join(stageData.Shell, "\",\"")))
		}

		if len(stageData.Cmd) > 0 {
			dockerfile.WriteString(fmt.Sprintf("CMD [\"%s\"]\n", strings.Join(stageData.Cmd, "\",\"")))
		}
		if len(stageData.Entrypoint) > 0 {
			dockerfile.WriteString(fmt.Sprintf("ENTRYPOINT [\"%s\"]\n", strings.Join(stageData.Entrypoint, "\",\"")))
		}
		dockerfile.WriteString("\n") // Add an empty line between stages
	}

	return dockerfile.String(), nil
}
