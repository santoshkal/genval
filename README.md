# GenVal

**GenVal** is a tool written in Golang, that generates a Dockerfile based on the input provided by the user in a `yaml` file. The tool encourages user to follow Dockerfile best-practices for Security and optimizating Docker images. Once the Dockerfile is generated. The tool further validates the generated Dockerfile using the rego framework against same Dockerfile best-practices. Once validated the tool prints helpful messages about the validation phase.

# TODO

- Update the Yaml parser to include all the Dockerfile instructions
- Update error messages for Validation steps
- Update messages for succussful validation of the Dockerfile


C, Go, Nodejs, Python, Rust

