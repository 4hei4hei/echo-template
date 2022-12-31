FROM golang:1.18.4

WORKDIR /app

COPY ["./", "/app/"]
RUN go mod tidy
CMD ["go", "run", "server.go"]
