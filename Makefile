build:
	go build .

run: main.go
	nodemon -e go,html --exec "go run" . --signal SIGTERM

deploy:
	flyctl deploy
