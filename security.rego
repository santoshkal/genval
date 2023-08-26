    package dockerfile_validation



allow{
	input[_].cmd == "from"
	val := split(input[_].value, "/")

	"cgr.dev" == val[0]
}

