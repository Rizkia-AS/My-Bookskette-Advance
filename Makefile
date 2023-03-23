pg:
	podman run --name DBbelajar -p 5432:5432 -e POSTGRES_USER=rizkia -e POSTGRES_PASSWORD=21204444 -d docker.io/library/postgres:15.2-alpine

createdb:
	podman exec -it DBbelajar createdb --username=rizkia --owner=rizkia bookskette

deletedb:
	podman exec -it DBbelajar dropdb --username=rizkia bookskette

miginit:
	migrate create -ext sql -dir db/migration -seq init_schema

migupdate:
	migrate create -ext sql -dir ./migrations update-schema

mignew:
	migrate create -dir db/migration -ext sql -seq new_db

migup:
	migrate -path db/migration -database "postgresql://rizkia:21204444@localhost:5432/bookskette?sslmode=disable" -verbose up

migdown:
	migrate -path db/migration -database "postgresql://rizkia:21204444@localhost:5432/bookskette?sslmode=disable" -verbose down

migforce:
	migrate -path db/migration -database "postgresql://rizkia:21204444@localhost:5432/bookskette?sslmode=disable" -verbose force 2

sqlcwipe:
	sqlc wipe

sqlcinit:
	sqlc init

sqlc:
	sqlc generate

test:
	go test -v

.PHONY: pg createdb deletedb miginit migupdate mignew migup migdown migforce sqlcwipe sqlcinit sqlc test