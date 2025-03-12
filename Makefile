.PHONY: clean

all: ldn-receiver ldn-sender

ldn-receiver: src/receiver/*.go
	go build -o ldn-receiver src/receiver/*.go

ldn-sender: src/sender/*.go 
	go build -o ldn-sender src/sender/*.go

clean:
	rm ldn-receiver ldn-sender