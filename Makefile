test:
	go test ./pkg/...

fmt:
	go fmt ./cmd/... ./pkg/...

vet:
	go vet ./cmd/... ./pkg/...

example:
	go run cmd/types-generator/main.go > pkg/example/generated-types.go
	@make fmt
