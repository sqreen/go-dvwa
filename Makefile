GO     ?= go
DOCKER ?= docker

.PHONY: all
all: dvwa

dvwa: sqreen-instrumentation-tool
	$(GO) build -v -o dvwa -a -toolexec $(PWD)/sqreen-instrumentation-tool ./server
help += server

sqreen-instrumentation-tool:
	$(GO) build -o sqreen-instrumentation-tool -v github.com/sqreen/go-agent/sdk/sqreen-instrumentation-tool
help += sqreen-instrumentation-tool

.PHONY: clean
clean:
	rm dvwa sqreen-instrumentation-tool

.PHONY: test
test:
	$(GO) test ./...

.PHONY: image
image:
	$(DOCKER) build -t go-dvwa .

.PHONY: run
run: image
	@$(DOCKER) run -it -e SQREEN_TOKEN="$(SQREEN_TOKEN)" -e SQREEN_APP_NAME="$(SQREEN_APP_NAME)" -p 8080:8080 go-dvwa