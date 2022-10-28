include .env
export $(shell sed 's/=.*//' .env)

test:
	@migrate -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:4011/${POSTGRES_DB}?sslmode=disable -path ./db/migrations down -all
	@migrate -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:4011/${POSTGRES_DB}?sslmode=disable -path ./db/migrations up
	@go clean -testcache && go test --race --cover ./...

.PHONY: migrate
make_migration:
	@migrate create -ext sql -dir db/migrations -seq $(name)

gen_swagger:
	@swag init --parseDependency

start_project:
	@docker-compose up -d

run_local_migrations:
	@migrate -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:4011/${POSTGRES_DB}?sslmode=disable -path ./db/migrations up

drop_local_migrations:
	@migrate -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:4011/${POSTGRES_DB}?sslmode=disable -path ./db/migrations down -all
