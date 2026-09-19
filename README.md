# Go-Zero Backend Starter

A production-oriented starter built on go-zero. One process exposes a generated-style REST API and optionally runs a Telegram long-polling transport. Both transports call the same application service.

## Stack

- Go 1.25+
- go-zero REST, config, logging, middleware and `goctl` API definition
- `go-telegram/bot`
- PostgreSQL via `pgxpool`
- `golang-migrate` migrations
- Docker Compose and GitHub Actions

## Layout

```text
api/starter.api          API source for goctl
cmd/api                  composition root
etc                      go-zero configuration
internal/handler         HTTP transport
internal/logic           use-case orchestration
internal/domain/note     framework-independent domain and repository
internal/transport       Telegram adapter
migrations               PostgreSQL migrations
```

## Quick start

```bash
cp .env.example .env
docker compose up --build
```

The bot is disabled when `TELEGRAM_BOT_TOKEN` is empty.

```bash
curl http://localhost:8080/healthz

curl -X POST http://localhost:8080/api/v1/notes \
  -H 'Content-Type: application/json' \
  -d '{"text":"first note"}'

curl 'http://localhost:8080/api/v1/notes?limit=20'
```

Telegram commands:

- `/start`
- `/add something to remember`
- `/notes`

## Development

Install goctl:

```bash
go install github.com/zeromicro/go-zero/tools/goctl@v1.10.2
```

After changing `api/starter.api`, regenerate the HTTP skeleton and then reapply intentional adapter customizations:

```bash
make generate
```

Run dependencies and the service:

```bash
docker compose up -d postgres
docker compose run --rm migrate
export DATABASE_URL='postgres://app:app@localhost:5432/app?sslmode=disable'
go mod tidy
go run ./cmd/api -f etc/starter-api.yaml
```

## What go-zero adds

Compared with the framework-free starter, go-zero owns server bootstrap, configuration mapping, request parsing, logging, recovery, timeouts and the `.api` code-generation workflow. Business rules and Telegram code stay outside the framework so they remain easy to test.

## Checks

```bash
make fmt
make vet
make test
```

## License

MIT
