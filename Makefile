start-dev:
	go run cmd/api/main.go

swag-generate:
	swag i --generalInfo cmd/api/main.go

swag-format:
	swag fmt