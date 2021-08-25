build:
	CGO_ENABLED=0 go build -mod=vendor -a -tags netgo -ldflags '-w' -o gcsb

test:
	go test -v --tags=json1 -mod=vendor -coverprofile=coverage.txt -covermode=atomic ./pkg/...

benchmark:
	go test -run ^$ -bench . -benchmem ./...