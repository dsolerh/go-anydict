build:
	go build -v ./...

test:
	go test -v ./...

bench:
	go test -benchmem -bench ^Benchmark githug.com/dsolerh/go-anydict/bench

coverage:
	go test -v -coverprofile=cover.out -covermode=atomic ./...
	go tool cover -html=cover.out -o cover.html

# tools
tools:
	go install github.com/psampaz/go-mod-outdated@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint:
	golangci-lint run --timeout 60s --max-same-issues 50 ./...
lint-fix:
	golangci-lint run --timeout 60s --max-same-issues 50 --fix ./...

audit: tools
	go list -json -m all | nancy sleuth

outdated:
	go install github.com/psampaz/go-mod-outdated@latest
	go list -u -m -json all | go-mod-outdated -update -direct

