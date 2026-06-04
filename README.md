# chirpy
To run the server correctly:

sqlc generate
go build ./...
goose -dir /home/ambujse/chirpy/chirpy/sql/schema postgres "postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable" up
DB_URL="postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable" PLATFORM=dev go run .

2026/06/04 14:33:58 goose: no migrations to run. current version: 3