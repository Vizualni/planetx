test:
	go test -count=10 ./...
build:
	go build -o build/planetx ./cmd/planetx
