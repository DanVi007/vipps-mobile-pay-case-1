# Backend

## Directory tree
```bash
.
|-- cmd
|-- constants
|-- handlers
`-- structs
```

## Entry point
cmd/main.go

## Dependencies
https://github.com/k3a/html2text

## Endpoints
### ../topic_counter?topic={value}
The topic param needs to be set other wise a bad request will be returned.
Returns amount of time topic is found from an article.
