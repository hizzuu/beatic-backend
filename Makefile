PJT_NAME = $(notdir $(PWD))
NET = app
SVC = api
DB_SVC = db
DB_NAME = app_dev

## Container up
.PHONY: up
up: down
	docker compose up -d

## Container down
.PHONY: down
down:
	docker compose down

## Generate gqlgen
.PHONY: gen-gql
gen-gql:
	 docker compose run --rm ${SVC} sh ./tools/gqlgen/generate.sh

# Cotainer attach
.PHONY: attach
attach:
	docker exec -it $(SVC) sh

# Generate migration $(name)
.PHONY: gen-migration
gen-migration:
	docker run --rm -v $(PWD)/build/mysql/migrations:/migrations migrate/migrate \
		create -ext sql -dir ./migration -seq $(name)

# Run migration $(status)
.PHONY: run-migration
run-migrate:
	docker run --rm -v $(PWD)/build/mysql/migrations:/migrations --network app migrate/migrate \
		-path=/migrations/ -database "mysql://root:password@tcp($(DB_SVC):3306)/$(DB_NAME)" $(status)
