
build:
	CGO_ENABLED=0 go build -a -o build/laptop-rotate-fix main.go

install:
	install build/laptop-rotate-fix /usr/local/bin/laptop-rotate-fix

.PHONY: build