.PHONY: test test-python test-go verify verify-go

PYTHON ?= python3

test: test-python test-go

test-python:
	PYTHONDONTWRITEBYTECODE=1 $(PYTHON) -m unittest discover -s tests -v

test-go:
	@if find . -name go.mod -not -path './.git/*' | grep -q .; then \
		./scripts/test-go-modules; \
	else \
		echo "Skipping Go tests: no go.mod files found"; \
	fi

verify: test-python verify-go

verify-go:
	@if find . -name go.mod -not -path './.git/*' | grep -q .; then \
		./scripts/verify-go; \
	else \
		echo "Skipping Go verification: no go.mod files found"; \
	fi
