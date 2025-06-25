createdb:
	docker exec -it 173d0f825da5 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it 173d0f825da5 dropdb simple_bank
	
dev:
	air

postgres:
	docker run --name simple_bank -e POSTGRES_PASSWORD=mypwd -e POSTGRES_USER=root -p 5432:5432 -d postgres:17.5-alpine

migrateup:
	migrate -path internal/db/sqlc/migrations/ -database "postgresql://root:mypwd@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/sqlc/migrations/ -database "postgresql://root:mypwd@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: createdb dropdb dev postgres migrateup migratedown sqlc