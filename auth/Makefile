-include .env

DB_URL = "postgresql://localhost:${POSTGRES_PORT}/postgres?user=${POSTGRES_USER}&password=${POSTGRES_PASSWORD}&sslmode=disable"

generate:
	oapi-codegen -generate gorilla -package api_client ./api/schema.yaml > ./api/generated/server.go
	oapi-codegen -generate types -package api_client ./api/schema.yaml > ./api/generated/types.go
	oapi-codegen -generate spec -package api_client ./api/schema.yaml > ./api/generated/swagger.go


migrate.up:
	migrate -database $(DB_URL) -path db/migrations up

migrate.down:
	migrate -database $(DB_URL) -path db/migrations down

migrate.create:
	migrate create -ext sql -dir db/migrations -seq $(name)