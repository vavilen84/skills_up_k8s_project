FROM golang:1.20-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY service1/ ./
COPY protobuf/ ./protobuf
RUN go build -o main .
CMD ["./main"]