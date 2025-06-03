build :
	@go build -o bin/app cmd/api/*.go

run : 
	@go run cmd/api/*.go
