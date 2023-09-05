    package dockerfile_validation


# Enforce a Base Image Prefix with Chainguard images:

untrusted_base_image[msg]{
	input[i].cmd == "from"
	val := split(input[i].value, "/")

	"cgr.dev" == val[0]
    msg := sprintf("Base image is prefixed with cgr.dev", [])
}


# # Avoid using 'latest' tag for Base Images:
latest_base_image[msg]{
    input[i].cmd == "from"
    val := split(input[i].value, ":")
    val[1] == "latest"
    msg := sprintf("Base image must not be tagged with 'latest'", [])
}

# Check for SemVer Tag for Base Images: NA as we are using `latest` for chainguard  images.



# # Detect Secrets in ENV Keys:
secrets_env = ["password", "secret", "key", "token"]

deny_secrets {
    input[i].Cmd == "env"
    val := input[i].Value
    contains(lower(val[_]), secrets_env[_])
}

multi_stage := true {
    input[_].Cmd == "copy"
    val := concat(" ", input[_].Flags)
    contains(lower(val), "--from=")
}
deny[msg] {
    multi_stage == false
    msg := sprintf("Not a multi-stage Dockerfile", [])
}

