.PHONY: install-tools lint test test-coverage ci

test:
	@echo ">  Running tests..."
	go test -v -race ./...

test-coverage:
	@echo ">  Running tests & coverage..."
	go test -v -race -coverprofile coverage.txt -covermode=atomic ./...

lint:
	@echo "  Running go vet..."
	go vet ./...
	@echo "  Running golint..."
	golint -set_exit_status=1 ./...

build: 
	go build -v ./...

install-tools:
	@echo ">  Installing tools..."
	go get -u golang.org/x/lint/golint

ci: install-tools lint test-coverage
