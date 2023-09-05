// validation/validation.go
package validation

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
		return fmt.Errorf("error reading rego policy: %v", err)
	}

	// Prepare Rego input data
	dockerfileInstructions := ParseDockerfileContent(dockerfileContent)

	jsonData, err := json.Marshal(dockerfileInstructions)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return nil
	}

	var commands []map[string]string
	err = json.Unmarshal([]byte(jsonData), &commands)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// fmt.Println("Commands:", jsonData)
	// policies := []string{"untrusted_base_image", "latest_base_image"}
	// for _, policy := range policies {
	// Create Rego for query and evaluation
	regoQuery := rego.New(
		// Rego rule package
		rego.Query("data.dockerfile_validation"),
		// rego.Query("data.dockerfile_validation.deny_root_user"),

		// rego policy file
		rego.Module("security.rego", string(regoPolicyCode)),

		// Dockerfile as input
		rego.Input(commands),
	)

	// Evaluate the Rego query
	rs, err := regoQuery.Eval(context.Background())
	if err != nil {
		log.Fatal("Error evaluating query:", err)
	}

	var extractedKeys []string
	for _, result := range rs {
		if len(result.Expressions) > 0 {
			// Iterate over the result
			for keys := range result.Expressions[0].Value.(map[string]interface{}) {
				extractedKeys = append(extractedKeys, keys)
			}
			for _, key := range extractedKeys {
				fmt.Printf("Policy: %s passed\n", key)
			}
		}
	}
	// fmt.Printf("Extracted keys: %v\n", key)

	// Get the number of policies evaluated by regoQuery
	fmt.Printf("Number of policies evaluated by regoQuery: %v\n", len(rs))

	if err != nil {
		return fmt.Errorf("error evaluating Rego: %v", err)
	}

	return nil
}
