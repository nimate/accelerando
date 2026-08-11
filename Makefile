.PHONY: test test-python test-go verify verify-go

PYTHON ?= python3

test: test-python test-go

test-python:
	PYTHONDONTWRITEBYTECODE=1 $(PYTHON) -m unittest discover -s tests -v

test-go:
	@if [ -f go.mod ]; then \
		go test ./...; \
	else \
		echo "Skipping Go tests: no go.mod"; \
	fi

verify: test-python verify-go

verify-go:
	@if [ -f go.mod ]; then \
		./scripts/verify-go; \
	else \
		echo "Skipping Go verification: no go.mod"; \
	fi
