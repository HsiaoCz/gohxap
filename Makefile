build:
	@go build -o bin/gohxap main.go

run: build
	@./bin/gohxap

gen:
	@templ generate

test:
	@go test -v ./...