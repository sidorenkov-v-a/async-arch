-include .env

generate:
	oapi-codegen -generate gorilla -package api ./api/schema.yaml > ./api/generated/server.go
	oapi-codegen -generate types -package api ./api/schema.yaml > ./api/generated/types.go

migrate.up:
	migrate -database ${DB_URL} -path db/migrations up

migrate.down:
	migrate -database ${DB_URL} -path db/migrations down