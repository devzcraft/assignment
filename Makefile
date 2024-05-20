run:
	go run cmd/main.go -config=.env
test:
	go test -race -v ./...
cover:
	go test -coverprofile=test.coverage.tmp ./... && go tool cover -func test.coverage.tmp
