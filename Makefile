build:
	go build . -o knave-tools

run: main.go
	nodemon -e go,html --exec "go run" . --signal SIGTERM
