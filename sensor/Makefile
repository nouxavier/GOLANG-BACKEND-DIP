.PHONY: migrate-create
# Generate a migration file in specified service directory
# first install with 'go get github.com/golang-migrate/migrate'
migrate-create:
	@migrate create -ext sql -dir db/migrations/$(SERVICE) -format unix $(NAME)


.PHONY: run-docker-postgres
#Run postgres
run-docker-postgres:
	@docker-compose up postgres
	