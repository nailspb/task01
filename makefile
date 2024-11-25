DB_USER:=user
DB_PASSWORD:=super_puper_user_password
DB_SERVER:=srv2.spartatn.ru
DB_PORT:=5430
DB_DATABASE:=tests

DB_DSN := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_SERVER):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable
MIGRATE := migrate -path ./migrations -database $(DB_DSN)
SEED := migrate -path ./seedes -database $(DB_DSN)

migrate-new:
	@migrate create -ext sql -dir ./migrations $(word 2, $(MAKECMDGOALS))

# Применение миграций
migrate:
	@$(MIGRATE) up

# Откат миграций
migrate-down:
	@$(MIGRATE) down

seed-up:
	@$(SEED) up

# Откат миграций
seed-down:
	@$(SEED) down

# для удобства добавим команду run, которая будет запускать наше приложение
run:
	@go run ./cmd/app/main.go

gen:
	@oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
	@oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go

lint:
	@golangci-lint run --out-format=colored-line-number
