/*
Gebug is a tool that makes debugging of Dockerized Go applications super easy by enabling Debugger and Hot-Reload features, seamlessly.


The implementation is based on utilizing Docker and Docker-Compose to manage the debugging environment.
In order to get a consistent and clean environment, the application build is taking place inside a Docker container.
This make the debugging process agnostic to the host's libraries and tools versions. For instance, you can choose to work a specific
version of Go on your host, but use another one when debugging your application.


How it works

During the `init` process, the user sets the desired configuration to the project including
Docker base, environment variables and expose ports.
When the user runs `start` (unless explicitly asked) a `Dockerfile` and `docker-compose.yml` with the relevant configuration
including one-direction source files synchronization between the host and the runtime container and auto-build on each change detected.

Configurations



	output_binary		- output binary artifact inside the runtime container (default: "/app")
	build_command		- build command inside the runtime container (default: "go build -o {{.output_binary}}")
	run_command		- run command, probably most of the time will just be the binary artifact path (default: "{.output_binary}}")
	runtime_image 		- base Docker image for the runtime container (default: "golang:1.18")
	debugger_enabled 	- whether to enable delve debugger inside the container or just use hot-reload (default: false)
	debugger_port 		- delve debugger listen port, relevant only if debugger_enabled was set (default: 40000)
	expose_ports 		- list of ports to expose inside the container (default: [])
	networks 		- list of docker external networks to join. if no network is selected, a new one will be created (default: [])
	environment 		- list of environment variables to be set inside the container (default: [])


Notes & Tips

	- Note you can reference other configuration fields.
	- When enabling Debugger -gcflags="all=-N -l" will be appended to the build command to stop compiler optimization and symbol removing.
	- No need to add the delve debugger listen port as it will be auto-added
	- Expose ports use the same syntax as docker-compose for mapping between host and container ports (e.g: "8080:8080")
	- Environment variables syntax: FOO=BAR or just FOO which will take the variable FOO from host and set it with its value

*/
package main
