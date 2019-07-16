all:
	go run cmd/webapp/*.go

run-db:
	docker-compose up -d

run-migrate-db:
	goose -dir deployments/databases/migration/ mysql "root:test@/example?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true" up

generate-openapi:
	goas --module-path . --main-file-path cmd/webapp/main.go --output api/openapi-spec/webapp-api.json
