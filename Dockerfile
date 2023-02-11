# IPBuf API Server Dockerfile

FROM golang:1.20

# Get the source code
WORKDIR /go/src/app
COPY . .

# Install dependencies
RUN go get -d -v ./...
RUN go install -v ./...

# Build and run the binary
RUN go build -o ipbuf .
CMD ["./ipbuf"]

# Expose the port
EXPOSE 8080
