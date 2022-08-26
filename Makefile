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
