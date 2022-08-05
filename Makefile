build:
	go build main.go knave.go -o knave-tools

run: main.go
	nodemon -e go,html --exec "go run" main.go knave.go --signal SIGTERM
