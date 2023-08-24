// validation/validation.go
package validation

import (
	"context"
	"encoding/json"
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

	jsonData, err := json.Marshal(dockerfileInstructions)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return nil
	}

	inputData := map[string]interface{}{
		"input": string(jsonData),
	}

	// fmt.Printf("Input data: %v\n", inputData)

	// policies := []string{"untrusted_base_image", "latest_base_image"}
	// for _, policy := range policies {
	// Create Rego for query and evaluation
	regoQuery := rego.New(
		// Rego rule package
		rego.Query("data.dockerfile_validation"),
		// rego.Query("data.dockerfile_validation.untrusted_base_image"),

		// rego policy filea
		rego.Module("./policy/security.rego", string(regoPolicyCode)),
		// Dockerfile as input
		rego.Input(inputData),
	)

	// Evaluate the Rego query
	rs, err := regoQuery.Eval(context.Background())

	// Get the number of policies evaluated by regoQuery
	fmt.Printf("Number of policies evaluated by regoQuery: %v\n", len(rs))
	// fmt.Printf("Number of policies evaluated by regoQuery: %v\n", rs[0].Expressions[0].Text)
	policyNames := []string{
		"untrusted_base_image",
		"latest_base_image",
		// Add more policy names as needed...
	}

	for i, r := range rs {
		for j, expr := range r.Expressions {
			if j < len(policyNames) {
				policyName := policyNames[j]
				allowed := expr.Value.([]interface{})
				fmt.Printf("Policy %d (%s): %v\n", i+1, policyName, allowed)
			}
		}
	}

	return nil
}
