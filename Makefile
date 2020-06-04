gopath 	= 	$(shell go env GOPATH)

install:
	- @echo "Installing app..."
	- go build -o $(gopath)/bin/goapp -i cmd/goapp/main.go
	- @echo "Installed"

build:
	- go build -o bin/goapp cmd/goapp/main.go

.PHONY: install build