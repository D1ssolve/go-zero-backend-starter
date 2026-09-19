.PHONY: run test fmt vet tidy generate up down

run:
	go run ./cmd/api -f etc/starter-api.yaml

test:
	go test -race -cover ./...

fmt:
	gofmt -w .

vet:
	go vet ./...

tidy:
	go mod tidy

generate:
	goctl api go -api api/starter.api -dir . --style=go_zero

up:
	docker compose up --build

down:
	docker compose down
