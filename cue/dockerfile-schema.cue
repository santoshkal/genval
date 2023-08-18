package docklerfile

dockerfile: [Dockerstage]

#Dockerstage: {
	stage?:  int
	from:    *string
	workdir: *string
	copy?:   *string
	run?: [...string]
	entrypoint?: [...string]
}
