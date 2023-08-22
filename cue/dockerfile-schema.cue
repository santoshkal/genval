package main

#Dockerfile:[...Dockerfile] 

#Dockerfile{
		stage: int
	from: string
	copy: [...string]
	workdir: string
	copy: string
	run: [...string]
	entrypoint: [...string]
    arg: string
	env: string
	expose: int
	volume: string
	user: string
	cmd: [...string]
	onbuild: string
	healthcheck: string
	shell: [...string]
	stopsignal: string
	label: string
	maintainer: string
}