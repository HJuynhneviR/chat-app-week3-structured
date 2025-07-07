# Chat App - Week 3 (Structured Version)

This version is split into folders for better readability and maintainability.

## Structure
- `controllers/` – WebSocket & HTTP handlers
- `models/` – Structs for Message & Client
- `middlewares/` – Rate limiter
- `utils/` – Redis initialization
- `main.go` – Bootstraps the server

## Run

```bash
go mod tidy
go run main.go
```

Ensure Redis is running on localhost:6379.
