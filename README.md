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

## Azure Function App Deployment

### Build

```bash
GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -trimpath main.go
```

### Test locally

```bash
func start
```

### Deploy

#### Use Visual Studio Code

Follow the instructions:
[Quickstart: Create a Go or Rust function in Azure using Visual Studio Code](https://learn.microsoft.com/en-us/azure/azure-functions/create-first-function-vs-code-other?tabs=go%2Cwindows#deploy-the-project-to-azure)

#### Use Azure CLI

```bash
func azure functionapp publish ipbuf
```
