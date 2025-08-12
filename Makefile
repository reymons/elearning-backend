build:
	go build -o main main.go

run:
	go run main.go

migrate:
	go run ./scripts/migrate.go
