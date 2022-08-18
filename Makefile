.PHONY: test

test:
	go test -v -race -coverprofile /dev/null ./...
