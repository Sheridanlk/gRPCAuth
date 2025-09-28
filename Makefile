.SILENT:

CONFIG_PATH = D:/gRPCSSO/sso/config/local.yaml
export CONFIG_PATH


set-path:
	@echo "CONFIG_PATH is $$CONFIG_PATH"

run: set-path
	go run ./cmd/sso/ .

create-migration:
	migrate create -ext sql -dir ./migrations -seq $(NAME)

migrate:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/sso?sslmode=disable' up

migrate-test:
	migrate -path ./tests/migrations -database 'postgres://postgres:postgres@localhost:5432/sso?sslmode=disable&x-migrations-table=migrations_test' up

migrate-down:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/sso?sslmode=disable' down
