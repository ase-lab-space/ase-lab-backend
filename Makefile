.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	go run .

.PHONY: gen
gen:
	go generate ./...
