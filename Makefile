include .env
export $(shell sed 's/=.*//' .env)

test:
	@go clean -testcache && go test --race --cover ./...

.PHONY: migrate
migrate:
	@migrate create -ext sql -dir db/migrations -seq $(name)