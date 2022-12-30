postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=passw0rd -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root expenses

dropdb:
	docker exec -it postgres12 dropdb expenses

migrateup:
	migrate -path db/migrations -database "postgresql://root:passw0rd@localhost:5432/expenses?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:passw0rd@localhost:5432/expenses?sslmode=disable" -verbose down

