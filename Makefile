all: clean build


test:
	@go test -race ./...

go_imports:
	@go install golang.org/x/tools/cmd/goimports@v0.22.0
	@goimports -w .
	@go mod tidy -compat=1.17

go_vet: go_imports
	@go vet ./...

go_lint: go_vet
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1
	@golangci-lint run

build: clean
	@go build -o bin/demo demo/main.go

clean:
	@rm -rf bin/demo

docker-build-and-push: clean
	@docker buildx create --name smsaero_golang --use || docker buildx use smsaero_golang
	@docker buildx build --platform linux/amd64,linux/arm64 -t 'smsaero/smsaero_golang:latest' . -f Dockerfile --push
	@docker buildx rm smsaero_golang
