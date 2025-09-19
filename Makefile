# generate code: mocks, openapi, etc
.PHONY: generate
generate:
	@go generate ./...

# run the app (suitable for development)
.PHONY: dev
dev:
	@go run cmd/estimates-api/main.go

# run unit tests
.PHONY: test
test:
	@go test ./...
