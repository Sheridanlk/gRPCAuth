.SILENT:

set-path:
	set "CONFIG_PATH=D:\gRPCSSO\sso\config\local.yaml" 

run: set-path
	go run ./cmd/sso/ .

create-migration:
	migrate create -ext sql -dir ./migrations -seq $(NAME)

migrate:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/sso?sslmode=disable' up

migrate-down:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/sso?sslmode=disable' down
