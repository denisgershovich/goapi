# Go API - Run Instructions

## Prerequisites
- Go 1.21+

## Quick Start
```powershell
# (optional) ensure deps
go mod tidy

# create env in root folder
.env

# run
go run ./cmd
```

## Configure Environment
Create or edit `.env` in the project root:
```dotenv
PORT=8080
DB_PATH=./data/app.db
```
Notes:
- Relative `DB_PATH` is relative to the project root.
- If not set, defaults to `./data/app.db`.
- `.env` is read on startup.

## First Run Behavior
- Creates SQLite file at `data/app.db` (or at `DB_PATH`).
- Applies SQL migrations from `internal/migrations`.

## Verify
```powershell
curl http://localhost:8080/health   # expect: ok
curl http://localhost:8080/api
```

## Project Structure
- `cmd/main.go`: entrypoint, routes
- `internal/config/`: `.env` loader
- `internal/db/`: SQLite + migrations
- `internal/migrations/`: SQL files
- `internal/handlers/`: HTTP handlers
- `internal/middlewares/`: logger
