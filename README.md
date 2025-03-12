# go-ldn-inbox

A Go implementation of an [LDN](https://www.w3.org/TR/ldn/) Receiver.

## Build the server

```
make
```

## Run the server

```
go-ldn-inbox -help
```

The service is now available at http://localhost:3333/inbox/

## POST Some Data

```
curl -X POST -H 'Content-Type: application/ld+json' --data-binary '@examples/demo.jsonld' http://localhost:3333/inbox/
```
## Optional

### Add some HTTP headers to the inbox

```
touch inbox/.meta
```

in the `.meta` file store as JSON key value pairs

## Help

```
Usage of ./go-ldn-inbox:
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