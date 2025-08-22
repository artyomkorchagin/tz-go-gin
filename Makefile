include .env

build:
	docker compose build

up:
	docker compose up

down:
	docker compose down

restart: down up

db-up:
	@goose -dir migrations postgres "user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=${DB_SSLMODE}" up

db-down:
	@goose -dir migrations postgres "user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=${DB_SSLMODE}" down

clean:
	docker compose down -v --rmi all