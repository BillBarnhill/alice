ENV_GOPATH ?= $(shell go env GOPATH)
GOTOOL_GOPATH ?= $(shell echo "$(ENV_GOPATH)" | awk '{n=split($$1, a, ":"); print a[n]}')
GOLINT ?= $(GOTOOL_GOPATH)/bin/golangci-lint

# lint
$(GOLINT):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOTOOL_GOPATH)/bin v1.62.0

PHONY += lint
lint: $(GOLINT)
	$(GOLINT) run --timeout 5m

