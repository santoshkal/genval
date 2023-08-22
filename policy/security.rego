    package dockerfile_validation

# evaluate_package imports and evaluates all rules within the package
# allow {
#    not untrusted_base_image
#    not latest_base_image
#     # deny_secrets
# }

# Enforce a Base Image Prefix with Chainguard images:
# allowed_images := ["cgr.dev"]

untrusted_base_image[msg]{
	input[i].Cmd == "from"
	val := split(input[i].Value[_], "/")

	"cgr.dev" == val[0]
    msg := sprintf("Base image is prefixed with cgr.dev", [])
}


# # Avoid using 'latest' tag for Base Images:
latest_base_image[msg]{
    input[i].Cmd == "from"
    val := split(input[i].Value[0], ":")
    val[1] != "latest"
    msg := sprintf("Base image must not be tagged with 'latest'", [])
}

# Check for SemVer Tag for Base Images:



# # Detect Secrets in ENV Keys:
# secrets_env = ["password", "secret", "key", "token"]

# deny_secrets {
#     input[i].Cmd == "env"
#     val := input[i].Value
#     contains(lower(val[_]), secrets_env[_])
# }

# multi_stage := true {
#     input[_].Cmd == "copy"
#     val := concat(" ", input[_].Flags)
#     contains(lower(val), "--from=")
# }
# deny[msg] {
#     multi_stage == false
#     msg := sprintf("Not a multi-stage Dockerfile", [])
# }

