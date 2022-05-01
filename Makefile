PWD=$(shell pwd)
def:
	docker run --rm -v ${PWD}:/app -w /app golangci/golangci-lint:v1.45.2 golangci-lint run -v --out-format=github-actions --path-prefix=. --timeout=5m

lint:
	golangci-lint run
