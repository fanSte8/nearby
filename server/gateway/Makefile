.PHONY: run
run:
	go run ./cmd/api/

.PHONY: audit
audit:
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Running tests...'
	go test -vet=off ./...

.PHONY: build
build:
	@echo 'Building...'
	go build -ldflags='-s' -o=./bin ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux_amd64/ ./cmd/api
