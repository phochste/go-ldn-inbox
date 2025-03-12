.PHONY: clean

all: ldn-receiver

ldn-receiver: receiver.go logger.go
	go build -o ldn-receiver *.go

clean:
	rm ldn-receiver