.SILENT:

CONFIG_PATH = C:\Users\sasha\Documents\go\gPRC\gRPCAuth\config\config.yaml
export CONFIG_PATH

DB_URL_DOCKER = postgres://postgres:postgres@auth-db:5432/auth?sslmode=disable
DB_URL_LOCAL = postgres://postgres:1234@localhost:5432/auth?sslmode=disable

set-path:
	@echo "CONFIG_PATH is $$CONFIG_PATH"

run-local: set-path
	go run ./cmd/sso/ .

run-docker: 
	docker-compose up --build

create-migration:
	migrate create -ext sql -dir ./migrations -seq $(NAME)

migrate-up-local:
	migrate -path ./migrations -database $(DB_URL_LOCAL) up

migrate-down-local:
	migrate -path ./migrations -database $(DB_URL_LOCAL) down 1

migrate-up-docker:
	migrate -path ./migrations -database $(DB_URL_DOCKER) up

migrate-down-docker:
	migrate -path ./migrations -database $(DB_URL_DOCKER) down 1

migrate-test:
	migrate -path ./tests/migrations -database 'postgres://postgres:1234@localhost:5432/auth?sslmode=disable&x-migrations-table=migrations_test' up
