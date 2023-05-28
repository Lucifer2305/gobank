build:
	@go build -o bin/gobank

ru: 
	@go test -v ./... 