FROM golang:1.23-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Install git and certs
RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY . .

# Build (this will auto-fetch go modules)
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]



