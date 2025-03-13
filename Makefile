.PHONY: install clean

all: ldn-receiver ldn-sender ldn-consumer

ldn-receiver: src/receiver/*.go
	go build -o ldn-receiver src/receiver/*.go

ldn-sender: src/sender/*.go 
	go build -o ldn-sender src/sender/*.go

ldn-consumer: src/consumer/*.go
	go build -o ldn-consumer src/consumer/*.go

install:
	go install ./src/ldn-consumer
	go install ./src/ldn-receiver
	go install ./src/ldn-sender

clean:
	rm ldn-receiver ldn-sender ldn-consumer