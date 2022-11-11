PWD=$(shell pwd)
default: lint test

dockerLint:
	docker run --rm -v ${PWD}:/app -w /app golangci/golangci-lint:v1.50.1 golangci-lint run -v --out-format=github-actions --path-prefix=. --timeout=5m

lint:
	golangci-lint run

test:
	go test ./... -gcflags=all=-l
