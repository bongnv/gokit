.PHONY: install-tools lint test test-coverage ci

test: check-tools
	@echo ">  Running tests..."
	go test -v -race ./...

test-coverage: check-tools
	@echo ">  Running tests & coverage..."
	go test -v -race -coverprofile coverage.txt -covermode=atomic ./...

lint: check-tools
	@echo "  Running go vet..."
	@echo $$go_files
	go vet ./...
	@echo "  Running golint..."
	$(golintable_files) | xargs -n 1 golint -set_exit_status=1

build: check-tools
	go build -v ./...

check-tools:
	$(check_all_tools)

install-tools:
	@echo ">  Installing tools..."
	go get -u golang.org/x/lint/golint

ci: install-tools lint test-coverage

lint_excluded_regex := (^|/)vendor/
lint_excluded_regex := $(lint_excluded_regex)|(^|/).*\.pb\.go
lint_excluded_regex := $(lint_excluded_regex)|(^|/)z_.*\.go
golintable_files = find . -type f -name "*.go" | grep -vE "$(lint_excluded_regex)"

required_tools := go golint
check_tool = $(shell command -v $(1) > /dev/null 2>&1 || (echo "$(1) is not installed, please install it"; exit 1))
check_all_tools = $(foreach tool,$(required_tools),$(call check_tool,$(tool)))