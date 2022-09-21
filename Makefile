test:
	go test -coverprofile=c.out ./...

coverage: test
	go tool cover -html=c.out

watch:
	ginkgo watch ./... -failFast

lint:
	gofmt -s -w .
	goimports -w -local github.com/miniscruff/igloo .
	golangci-lint run ./...
