.PHONY: build
build:
				go build -v ./cmd/apiserver

.PHONY: test
test: 
				go test -v -race -timeout 30s  ./...

admin:
				go build -v ./internal/createadmin
.DEFAULT_GOAL := build