.PHONY: mod
mod:
	go get -u
	go mod tidy
	go mod verify

.PHONY: run
run:
	go run .

.PHONY: test
test:
	go test -v ./...
