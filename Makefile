include .env
export $(shell sed 's/=.*//' .env)

MIGRATE_FILE_DIR := migrations/files

clean_module:
	go mod tidy

run:
	go run app/main.go

run_test:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

compose_up:
	docker-compose down
	docker network remove pokedex-api_pokedex_network; true
	docker-compose up --force-recreate

compose_down:
	docker-compose down

migration_create:
	migrate create -ext sql -dir $(MIGRATE_FILE_DIR) $(NAME)

migration_up:
	go run migrations/app/main.go -type up

migration_down:
	go run migrations/app/main.go -type down -version $(VERSION)

migration_force:
	go run migrations/app/main.go -type force -version $(VERSION)