.PHONY: build clean tool lint help

all: build

review: cover-profile cover-html

build:
	@go build -v .

tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

clean:
	rm -rf flash-sale
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"

test:
	@go test -v ./...

run:
	@go run cmd/main.go

check-doc-reqs:
	$(foreach bin,$(DOC_EXECUTABLES),\
		$(if $(shell command -v $(bin) 2> /dev/null),$(info Found `$(bin)`),$(error Please install `$(bin)`)))

generate-doc: check-doc-reqs
	@echo "Generating swagger.yaml"
	@swagger generate spec -o ./docs/swagger.yaml  --scan-models

cover-profile:
	@go test -v -coverprofile cover.out ./...

cover-html:
	@go tool cover -html=cover.out -o cover.html