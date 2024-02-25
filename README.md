# Print Incoming Request

A simple app that prints incoming POST requests, created to experiment with the new mux.

Install

```bash
go install github.com/mxssl/pir@latest
```

Test

```bash
curl \
  --url 'localhost:9999' \
  --request POST \
  --data '{"key":"value"}'
```
