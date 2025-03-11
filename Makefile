all: main.go logger.go
	go build -o go-ldn-inbox *.go
