# assembly-visual-backend
This repo contains a Go (Gin + GORM) API backed by PostgreSQL.  
The project supports **Windows (CMD/PowerShell)** via a `.bat` script and **Linux/macOS** via a `Makefile`.  
Swagger UI is exposed at **`/swagger/index.html`** when docs are generated.

---

## 1) Prerequisites

- Go 1.21+
- PostgreSQL 15+ (or Docker for `postgres` + `pgadmin`)
- (Optional) `swag` CLI for generating Swagger docs:  
    ```bash
        go install github.com/swaggo/swag/cmd/swag@latest
    ```

## 2) Environment Variables
- Create a file named `.env` in the project root:
    ```bash
        POSTGRES_USER=postgres
        POSTGRES_PASSWORD=postgres
        POSTGRES_DB=assembly-visual

        # If backend runs outside Docker, talking to Postgres in Docker on the same machine
        POSTGRES_HOST=localhost

        # If backend runs inside the same docker-compose as Postgres
        # POSTGRES_HOST=postgres

        POSTGRES_PORT=5432
        POSTGRES_SSL=disable
        POSTGRES_TIMEZONE=Asia/Bangkok

        ENV=dev
        FE_URL=http://localhost:3000
        PORT=9090
        AUTO_MIGRATE=false
    ```
- **Tip**: Use localhost when your API runs outside Docker. Use postgres when your API is a service inside the same docker-compose network as the DB.

## 3) Run with Docker (optional but recommended)
    ```bash
        docker compose up -d
        # or
        docker-compose up -d
    ```
- This brings up:

    - `postgres` on `5432`
    - `pgadmin` on `http://localhost:5050` (if included in your compose)

- Connect pgAdmin → Host `postgres` (inside docker network) or `localhost` (from your host OS).

## 4) Windows (CMD/PowerShell) – use build.bat
**Commands:**

    ```bash
        build build             # go build -o main cmd\api\main.go
        build run               # go run cmd\api\main.go
        build test              # go test .\test\... -v
        build swagger-install   # go install swag CLI
        build swagger           # swag init -g cmd\api\main.go -o cmd\api\docs
        build clean             # remove build artifacts

    ```
Open **CMD** in the project root (same folder as build.bat).


## 5) Linux/macOS – use Makefile
**Commands:**

    ```bash
        make build
        make run
        make test
        make swagger-install
        make swagger
        make clean
    ```

## 6) Swagger (OpenAPI)
**Generate docs**
- Windows:
```bash

    build swagger-install
    build swagger

```
- Linux/macOS:
```bash

    make swagger-install
    make swagger

```
**Open Swagger UI**
```bash
http://localhost:9090/swagger/index.html
```

## 7) API Endpoints (current snapshot)
**Assuming base URL http://localhost:9090**

- GET /api/v2/users/ — list users

- POST /api/v2/users/signup — create user

- GET /api/v2/cats/ — list cats (demo)

- POST /api/v2/operations/add — add numbers

- POST /api/v2/operations/sub — subtract numbers

- POST /api/v2/operations/mul — multiply numbers

- POST /api/v2/operations/div — divide numbers

- GET /api/v2/oauth/google/callback — Google OAuth callback

- GET /swagger/*any — Swagger UI (/swagger/index.html)

## 8) Run directly without scripts (if needed)
```bash
# build
go build -o main cmd/api/main.go

# run
go run cmd/api/main.go

# test
go clean -testcache
go test ./test/... -v

```