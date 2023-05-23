build:
	@go build -o bin/booli-api-go

run: 	build	
	@./bin/booli-api-go
test: 
	@go test -v ./...
