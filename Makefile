.PHONY: lint

test:
	go test -v ./...

lint:
	golangci-lint run
