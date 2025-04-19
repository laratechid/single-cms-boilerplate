dev:
	air

start:
	go run cmd/api/main.go

test:
	cp .env internal/service && go test ./internal/service && rm internal/service/.env

swag-generate:
	swag i --generalInfo cmd/api/main.go

swag-format:
	swag fmt

mock-create:
	mockgen -source=$(source) -destination=$(basename $(source))_mock.go