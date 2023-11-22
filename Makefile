build:
	@go build -o ./tmp/main .

run: build
	@./tmp/main