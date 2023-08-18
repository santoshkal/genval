    package dockerfile_validation

# evaluate_package imports and evaluates all rules within the package
evaluate_package {
     untrusted_base_image
     latest_base_image
    # deny_secrets
}

# Enforce a Base Image Prefix with Chainguard images:
untrusted_base_image {
    input[i].Cmd == "from"
    val := split(input[i].Value, "/")
    val[0] == "c"  
}


# Avoid 'Latest' Tag for Base Images:
latest_base_image {
    input[i].Cmd == "from"
    val := split(input[i].Value[0], ":")
    not contains(lower(val[1]), "latest")
}

# # Detect Secrets in ENV Keys:
# secrets_env = ["password", "secret", "key", "token"]

# deny_secrets {
#     input[i].Cmd == "env"
#     val := input[i].Value
#     contains(lower(val[_]), secrets_env[_])
# }
