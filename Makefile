GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOMODDOWNLOAD=$(GOCMD) mod tidy

APP_ENTRYPOINT=cmd/main.go
BINARY_NAME=main
BUILD_FOLDER=build

.PHONY: all build clean deps

all: clean build

build:
	mkdir -p $(BUILD_FOLDER)
	cp -n .env $(BUILD_FOLDER)/.env
	$(GOBUILD) -o $(BUILD_FOLDER)/$(BINARY_NAME) -v $(APP_ENTRYPOINT)
deps:
	$(GOMODDOWNLOAD)
clean:
	$(GOCLEAN)
	rm -rf $(BUILD_FOLDER)/*
start:
	./$(BUILD_FOLDER)/$(BINARY_NAME) -c $(BUILD_FOLDER)/.env
run: clean build start