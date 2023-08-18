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
