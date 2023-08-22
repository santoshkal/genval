    package dockerfile_validation

# Write Rego policies for Dockerfile validation that will be used by the Golang script to validate Dockerfiles for checking if Image starts with `cgr.dev` and does not contain `latest` tag.

validate_dockerfile.rego

    package dockerfile_validation

    violation[msg] {
        input := parse_dockerfile(input.content)
        input.base_image.tag != "latest"
        not startswith(input.base_image.name, "cgr.dev")
        msg := sprintf("Dockerfile %v violates the policy", [input.path])
    }