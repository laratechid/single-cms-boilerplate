# === COLORS ===
yellowText=\033[0;33m
resetText=\033[0m

# === FUNCTIONS ===
define print-example
	@echo "$(yellowText)example command:$(resetText) make $(1)"
	@echo ""
endef

# === COMMANDS ===

develop:
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
	$(call print-example,mock-create source=internal/repository/user.go)
	mockgen -source=$(source) -destination=$(basename $(source))_mock.go

migration-create:
	$(call print-example,migration-create name=add_role_permission)
	migrate create -ext sql -dir ./migrations $(name)

migration-up:
	$(call print-example,migration-up dbUrl=postgres://postgres:root@localhost:5432/single_brand)
	migrate -path ./migrations -database "$(dbUrl)" up
