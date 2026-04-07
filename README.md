# Simple URL Shortener API

A production-ready URL shortener service built with Docker, **Go**, **Fiber**, and **SQLite**.

## Features
- **Shorten URLs**: Quickly generate a short link.
- **Redirects**: Fast redirects to the original long URL.
- **Stats**: Track the number of clicks for each generated short URL.
- **SQLite Database**: Persistent relational database out of the box.
- **Dockerized**: Easy deployment using Docker.

## Project Structure
We utilize a clean architecture for high maintainability:
- `cmd/server/main.go` - Application entry point
- `internal/handlers` - HTTP layer mapping requests to services
- `internal/services` - Core business logic and validations
- `internal/repositories` - Database and data persistence layer
- `internal/database` - DB initialization and configs
- `internal/models` - Domain models

## Run with Docker

1. Build the Docker image:
```bash
docker build -t simple-url-shortener .
```

2. Run the Docker container:
```bash
docker run -p 3000:3000 simple-url-shortener
```

## Running Locally

1. Install dependencies:
```bash
go mod tidy
```

2. Run server:
```bash
go run cmd/server/main.go
```

## Endpoints

### 1. Shorten URL
**POST** `/shorten`

**Payload:**
```json
{
  "long_url": "https://www.example.com/some/very/long/path"
}
```

**Response:** (201 Created)
```json
{
  "id": 1,
  "long_url": "https://www.example.com/some/very/long/path",
  "short_code": "kH7z9P",
  "clicks": 0,
  "created_at": "...",
  "updated_at": "..."
}
```

### 2. Redirect
**GET** `/:code`

Automatically redirects (with `301 Moved Permanently`) to the original long URL.

### 3. Retrieve Stats
**GET** `/stats/:code`

**Response:** (200 OK)
```json
{
  "id": 1,
  "long_url": "https://www.example.com/some/very/long/path",
  "short_code": "kH7z9P",
  "clicks": 5,
  "created_at": "...",
  "updated_at": "..."
}
```
