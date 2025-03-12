# go-ldn-inbox

A Go implementation of an [Linked Data Notifications](https://www.w3.org/TR/ldn/) Sender, Consumer and Receiver.

## Build the server

```
make
```

## Run the server

```
./ldn-receiver
```

The service is now available at http://localhost:3333/inbox/

## Send some data

```
./ldn-sender http://localhost:3333/inbox/ ./examples/demo.jsonld
```

## Read some data

```
./ldn-consumer http://localhost:3333/inbox/
http://localhost:3333/inbox/b09dee2e4dc5cecdf0951dab287a906d.jsonld
./ldn-consumer http://localhost:3333/inbox/b09dee2e4dc5cecdf0951dab287a906d.jsonld
...{json-ld output}...
```

## Optional

### Add some HTTP headers to the inbox

```
touch inbox/.meta
```

in the `.meta` file store as JSON key value pairs

## Help

```
Usage of ./ldn-receiver:
  -base string
        Base URL (default "http://localhost:3333")
  -host string
        Hostname (default "localhost")
  -inboxDir string
        Local path to your inbox (default "./inbox")
  -inboxPath string
        URL path to your inbox (default "/inbox/")
  -port int
        Port (default 3333)
  -public
        World readable inbox (default true)
  -schema string
        JSON schema to validate input
  -writable
        World appendable inbox (default true)
```

## Author

Patrick Hochstenbach 

MIT license