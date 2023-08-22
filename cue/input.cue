dockerfile: [{
	stage:   1
	from:    "golang:1.20-alpint"
	workdir: "/app"
	copy:    ". ."
	run: [
		"go mod download",
		"go build -o app",
	]
}, {
	stage:   2
	from:    "scratch"
	workdir: "/app"
	copy:    "/app/app ."
	entrypoint: [
		"./app",
	]
}]


# write a cuelang schema for the abpve yaml in cuelang
# https://cuelang.org/docs/tutorials/tour/schema/
# https://cuelang.org/docs/references/spec/#schema
# https://cuelang.org/docs/references/spec/#schema-definitions
# https://cuelang.org/docs/references/spec/#schema-attributes
# https://cuelang.org/docs/references/spec/#schema-attribute-values
# https://cuelang.org/docs/references/spec/#schema-attribute-values
# https://cuelang.org/docs/references/spec/#schema-attribute-values
# https://cuelang.org/docs/references/spec/#schema-attribute-values
# https://cuelang.org/docs/references/spec/#schema-attribute-values


# write a cuelang schema for the abpve yaml in cuelang and No links please

