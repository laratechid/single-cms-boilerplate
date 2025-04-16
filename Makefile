start-dev:
	go run cmd/api/main.go

start-test:
	cp .env internal/service && go test ./internal/service && rm internal/service/.env

swag-generate:
	swag i --generalInfo cmd/api/main.go

swag-format:
	swag fmt

