    package dockerfile_validation

allow {
   not untrusted_base_image
   not latest_base_image
}


untrusted_base_image{
	input[i].cmd == "from"
	val := split(input[i].value, "/")

	"cgr.dev" != val[0]
}


latest_base_image{
    input[i].cmd == "from"
    val := split(input[i].value, ":")
    val[1] == "latest"
}

