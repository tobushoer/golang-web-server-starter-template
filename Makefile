all:
	docker-compose up -d
	go run cmd/webapp/*.go

generate-openapi:
	goas --module-path . --main-file-path cmd/webapp/main.go --output api/openapi-spec/todo-v1.json
