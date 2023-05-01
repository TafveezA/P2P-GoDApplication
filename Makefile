hello:
	echo "Hello, World"
build:
	@go build -o bin/p2pgame
run:
	@./bin/p2pgame
test:
	go test -v ./ ..