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

.PHONY: fetch-oui-list
fetch-oui-list:
	wget "http://standards-oui.ieee.org/oui/oui.csv" -O "./assets/oui.csv"
