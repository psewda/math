BASE_DIR=$(shell pwd)
OUTPUT=$(BASE_DIR)/bin
APP=math
MAIN=$(BASE_DIR)/main.go

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64
	go build -ldflags="-s -w " -o $(OUTPUT)/$(APP) $(MAIN)

test:
	go test -v ./...
