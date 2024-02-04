APP_BINARY=shopsApp

postgres:
	docker run --name postgres -p 5402:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d postgres

createdb:
	docker exec -it shops_postgres_1 createdb --username=postgres --owner=postgres shops

dropdb:
	docker exec -it postgres dropdb shops

migrateup:
	migrate -path pkg/db/migrations -database "postgresql://postgres:password@localhost:5402/shops?sslmode=disable" -verbose up

migrateup1:
	migrate -path pkg/db/migrations -database "postgresql://postgres:password@localhost:5402/shops?sslmode=disable" -verbose up 1
	
migratedown:
	migrate -path pkg/db/migrations -database "postgresql://postgres:password@localhost:5402/shops?sslmode=disable" -verbose down

migratedown1:
	migrate -path pkg/db/migrations -database "postgresql://postgres:password@localhost:5402/shops?sslmode=disable" -verbose down 1

sqlc:
	docker run --rm -v ${CURDIR}:/src -w /src sqlc/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

buildgo:
	chdir . && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${APP_BINARY} ./cmd/app/

mock:
	mockgen -package mockdb -destination pkg/db/mock/store.go github.com/zura-t/go_delivery_system-shops/pkg/db/sqlc Store

.PHONY: postgres test sqlc createdb dropdb mock migratedown migrateup migratedown2 migrateup1 server build