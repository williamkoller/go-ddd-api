MIGRATIONS_DIR=./migrations
DB_URL_DEV=postgresql://postgres:postgres@localhost:5432/go_ddd?sslmode=disable
DB_URL_DOCKER=postgresql://postgres:postgres@db:5432/go_ddd?sslmode=disable

dev:
	docker-compose up --build

build:
	go build -o main ./cmd/main.go

run:
	go run ./cmd/main.go

test:
	go test ./internal/... ./pkg/...

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL_DEV)" up

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL_DEV)" down

migrate-up-docker:
	docker run --network=backend --rm -v $(PWD)/migrations:/migrations migrate/migrate -path=/migrations -database "$(DB_URL_DOCKER)" up

migrate-down-docker:
	docker run --network=backend --rm -v $(PWD)/migrations:/migrations migrate/migrate -path=/migrations -database "$(DB_URL_DOCKER)" down
