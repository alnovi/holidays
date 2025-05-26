.SILENT:
.DEFAULT-GOAL:= help

## help: справка
.PHONY: help
help:
	@echo 'HOLIDAYS'
	@echo ''
	@echo 'Usage:'
	@echo '  make <command>'
	@echo ''
	@echo 'The commands are:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## migration: новая миграция
.PHONY: migration
migration:
	@read -p "Enter migration name: " migration_name; \
	go tool goose -dir=./scripts/migrations create $$migration_name go

## swag: генерация документации
.PHONY: swag
swag:
	@go tool swag init -g ./cmd/server/main.go -p snakecase

## lint: статический анализ
.PHONY: lint
lint:
	go tool golangci-lint run ./...

## lint-fix: статический анализ (авто-исправление)
.PHONY: lint-fix
lint-fix:
	go tool golangci-lint run ./... --fix --timeout 650s
