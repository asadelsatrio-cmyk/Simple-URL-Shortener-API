FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install build dependencies for go-sqlite3
RUN apk add --no-cache build-base

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application with CGO enabled (required for sqlite3)
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o url-shortener ./cmd/server

FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/url-shortener .

EXPOSE 3000

CMD ["./url-shortener"]
