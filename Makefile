.PHONY: install clean

all: ldn-receiver ldn-sender ldn-consumer

ldn-receiver: src/ldn-receiver/*.go
	go build -o ldn-receiver src/ldn-receiver/*.go

ldn-sender: src/ldn-sender/*.go 
	go build -o ldn-sender src/ldn-sender/*.go

ldn-consumer: src/ldn-consumer/*.go
	go build -o ldn-consumer src/ldn-consumer/*.go

install:
	go install ./src/ldn-consumer
	go install ./src/ldn-receiver
	go install ./src/ldn-sender

clean:
	rm ldn-receiver ldn-sender ldn-consumer