run:
	go run cmd/main.go

sqlc:
	sqlc generate

.PHONY: run sqlc