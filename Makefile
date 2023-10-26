postgres:
	docker run --name postgres_mybank -p 5455:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres_mybank createdb --username=root --owner=root my_bank

dropdb:
	docker exec -it postgres_mybank dropdb my_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5455/my_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5455/my_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test