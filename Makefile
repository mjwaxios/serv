serv:	main.go
	CGO_ENABLED=0 go build

image:	serv
	podman build . -t serv

copy:
	skopeo copy containers-storage:localhost/serv:latest docker://localhost:5000/serv:$(VER)