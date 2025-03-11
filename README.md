# go-ldn-inbox

A Go implementation of an LDN Inbox

## Run the server

```
go run main.go logger.go
```

The service is now available at http://localhost:3333

## POST Some Data

```
curl -X POST -H 'Content-Type: application/ld+json' --data-binary '@examples/demo.jsonld' http://localhost:3333/inbox/
```
