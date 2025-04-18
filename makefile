include .env
export

MIGRATE=migrate -path db/migrations -database "$(DB_URL)"

.PHONY: migrate-up migrate-down migrate-force migrate-create

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

migrate-force:
	$(MIGRATE) force

migrate-version:
	$(MIGRATE) version

migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(NAME)
