.PHONY: build test vet

build:
	GOOS=darwin GOARCH=amd64 go1.9beta2 build -o twilioagg-osx-amd64 github.com/adamdecaf/twilioagg
	GOOS=linux GOARCH=amd64 go1.9beta2 build -o twilioagg-linux-amd64 github.com/adamdecaf/twilioagg

vet:
	go1.9beta2 vet .

test: vet
	go1.9beta2 test -v .
