    package dockerfile_validation

# allow {
#    not untrusted_base_image
#    not latest_base_image
# }


untrusted_base_image contains msg if {
	input[i].cmd == "from"
	val := split(input[i].value, "/")

	"cgr.dev" == val[0]
	msg := sprintf("Val: %v", [val[0]])
}

latest_base_image contains msg if {
	input[i].cmd == "from"
	val1 := split(input[i].value, ":")
	contains(val1[1], "latest")
	msg := sprintf("Val1: %v", [val1[1]])
}



# https://play.openpolicyagent.org/p/x9pwuimWKd