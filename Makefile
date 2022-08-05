build:
	go build main.go

run: main.go
	nodemon -e go,html --exec "go run" main.go --signal SIGTERM
