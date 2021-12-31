PROGRAM_NAME=myza

init:
	go mod init

tidy:
	go mod tidy

build:
	go build -o $(PROGRAM_NAME)

tests:
	go test ./...

mocks:
	mockery -all -dir internal/domain -output internal/repository/mocks