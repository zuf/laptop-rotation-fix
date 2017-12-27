
build:
	CGO_ENABLED=0 go build -a -o build/laptop-rotation-fix main.go

install:
	install build/laptop-rotation-fix /usr/local/bin/laptop-rotation-fix

.PHONY: build