// validation/validation.go
package validation

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/open-policy-agent/opa/rego"
)

// DockerfileInstruction represents a Dockerfile instruction with Cmd and Value.
type DockerfileInstruction struct {
	Cmd   string `json:"cmd"`
	Value string `json:"value"`
}

// ...

func ParseDockerfileContent(content string) []DockerfileInstruction {
	lines := strings.Split(content, "\n")
	var instructions []DockerfileInstruction

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		cmd := strings.ToLower(parts[0])
		value := strings.Join(parts[1:], " ")

		instructions = append(instructions, DockerfileInstruction{
			Cmd:   cmd,
			Value: value,
		})
	}

	return instructions
}

// ValidateDockerfileUsingRego validates a Dockerfile using Rego.
func ValidateDockerfileUsingRego(dockerfileContent string, regoPolicyPath string) error {
	// Read Rego policy code from file
	regoPolicyCode, err := os.ReadFile(regoPolicyPath)
	if err != nil {
		return fmt.Errorf("Error reading Rego policy: %v", err)
	}

	// Prepare Rego input data
	dockerfileInstructions := ParseDockerfileContent(dockerfileContent)

	inputData := map[string]interface{}{
		"input": dockerfileInstructions,
	}

	// Create Rego for query and evaluation
	regoQuery := rego.New(
		// Rego rule package
		rego.Query("data.dockerfile_validation.evaluate_package"),
		// rego policy file
		rego.Module("./policy/security.rego", string(regoPolicyCode)),
		// Dockerfile as input
		rego.Input(inputData),
	)

	// Evaluate the Rego query
	rs, err := regoQuery.Eval(context.Background())
	if err != nil {
		return fmt.Errorf("Error evaluating Rego: %v", err)
	}
	if rs.Allowed() == true {
		fmt.Printf("Dockerfile Valdation Succeeded with %v\n", rs.Allowed())
	} else {
		fmt.Printf("Dockerfile Validation failed %v\n", rs.Allowed())
	}
	return nil

}
