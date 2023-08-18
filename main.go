package main

import (
	"fmt"
	"log"
	"os"

	"github.com/santoshkal/genforce/generator"
	"github.com/santoshkal/genforce/validation"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <path-to-input-yaml> <output-dockerfile-path>")
		return
	}

	inputYAMLPath := os.Args[1]
	outputDockerfilePath := os.Args[2]

	// Read input YAML
	yamlData, err := os.ReadFile(inputYAMLPath)
	if err != nil {
		log.Fatal("Error reading input YAML file:", err)
	}

	// Step 1: Generate Dockerfile
	generatedDockerfileContent, err := generator.GenerateDockerfile(yamlData)
	if err != nil {
		log.Fatal("Error generating Dockerfile from YAML:", err)
	}

	// Save generated Dockerfile to the specified path
	err = os.WriteFile(outputDockerfilePath, []byte(generatedDockerfileContent), 0644)
	if err != nil {
		log.Fatal("Error writing generated Dockerfile:", err)
	}

	fmt.Printf("Generated Dockerfile saved to: %s\n", outputDockerfilePath)

	// Step 2: Validate Dockerfile using Rego
	// Update the path to your policy file
	policyFilePath := "./policy/security.rego"
	err = validation.ValidateDockerfileUsingRego(generatedDockerfileContent, policyFilePath)
	if err != nil {
		fmt.Println("Dockerfile validation failed:", err)
		return
	} else {
		fmt.Println("Dockerfile validation succeeded!")
	}
}
