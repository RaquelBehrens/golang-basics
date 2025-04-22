test all:
	go test -v ./...

coverage:
	go test -cover -coverprofile=coverage.out ./...

coverage-html: coverage
	go tool cover -html="coverage.out"

coverage-total: coverage
	go tool cover -func="coverage.out"
