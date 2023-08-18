    package dockerfile_validation


# Avoid 'Latest' Tag for Base Images:
latest_base_image {
    input[i].Cmd == "from"
    val := split(input[i].Value[0], ":")
    contains(lower(val[1]), "latest")
}