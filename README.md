# go-ldn-inbox

A Go implementation of an LDN Inbox

## Build the server

```
make
```

## Run the server

```
go-ldn-inbox -help
```

The service is now available at http://localhost:3333

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

## Author

Patrick Hochstenbach 

MIT license