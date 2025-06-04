build :
	@go build -o bin/app cmd/api/*.go

run : build
	@./bin/app
