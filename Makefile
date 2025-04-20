# === COLORS ===
yellowText=\033[0;33m
resetText=\033[0m

# === COMMANDS ===

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
	@echo "$(yellowText)example command:$(resetText) make mock-create source=internal/repository/user.go"
	@echo ""
	mockgen -source=$(source) -destination=$(basename $(source))_mock.go

migration-create:
	@echo "$(yellowText)example command:$(resetText) make migration-create name=add_role_permission"
	@echo ""
	migrate create -ext sql -dir ./migrations $(name)

migration-up:
	@echo "$(yellowText)example command:$(resetText) make migration-up dbUrl=postgres://postgres:root@localhost:5432/single_brand"
	@echo ""
	migrate -path ./migrations -database "$(dbUrl)" up
