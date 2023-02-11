# IPBuf

## Compile `.proto` files

```bash
protoc --go_out=proto proto/main.proto
```

## Run the server

### Docker

```bash
docker build -t ipbuf . && \
docker run -p 8080:8080 \
--name ipbuf-container ipbuf-image
```
