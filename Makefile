
build:
	CGO_ENABLED=0 go build -a -o build/laptop-rotate-fix main.go

.PHONY: build