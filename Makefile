PROGRAM_NAME=myza

init:
	go mod init

tidy:
	go clean -modcache
	go mod tidy

build:
	go build -o $(PROGRAM_NAME)

tests:
	go test ./...

mocks:
	mockery --all --dir domain --output domain/mocks