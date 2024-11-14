include .env
MIGRATIONS_PATH = cmd/migrations

.PHONY: build-posts
build-posts:
	@go build -o bin/posts cmd/posts/main.go

.PHONY: build-users
build-users:
	@go build -o bin/users cmd/users/main.go

.PHONY: run-users
run-users: build-users
	@./bin/users

.PHONY: run-posts
run-posts: build-posts
	@./bin/posts

.PHONY: gen-proto
gen-proto:
	@protoc --proto_path=pkg/proto \
		--go_out=pkg/proto/generated --go_opt=paths=source_relative \
		--go-grpc_out=pkg/proto/generated --go-grpc_opt=paths=source_relative \
		pkg/proto/*.proto

.PHONY: gen-migration
gen-migration:
	@name=$(name) ; \

	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(name)

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(POSTGRES_URI) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(POSTGRES_URI) down $(filter-out $@,$(MAKECMDGOALS))