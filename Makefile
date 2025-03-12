all: main.go logger.go
	go build -o ldn-receiver *.go
