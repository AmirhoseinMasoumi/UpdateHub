DB_URL=postgresql://root:secret@localhost:5432/update_hub?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root update_hub

dropdb:
	docker exec -it postgres dropdb update_hub

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate	

mock:
	mockgen -destination db/mock/store.go  github.com/AmirhoseinMasoumi/GoProjects/DeviceUpdateManager/db/sqlc Store

server:
	go run main.go

.PHONY: createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc mock server

