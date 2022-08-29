TOOLS=\
	github.com/rubenv/sql-migrate/...@v1.1.2 \
	github.com/cosmtrek/air@latest

.PHONY: all
all:

.PHONY: install-tools
install-tools:
	@for tool in ${TOOLS} ; do \
		echo "install $${tool}"; \
		go install $${tool}; \
	done

.PHONY: test
test:
	@go test -v --race --cover -coverprofile=coverage.txt -covermode=atomic ./...
	@go tool cover -html=coverage.txt -o coverage.html
