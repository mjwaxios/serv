serv:	main.go
	CGO_ENABLED=0 go build

image:	serv
	podman build . -t serv
