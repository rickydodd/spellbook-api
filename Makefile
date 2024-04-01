MAIN_PACKAGE_PATH := ./cmd/server
BINARY_NAME := api
PORT ?= 80

# ========================= #
# Quality control
# ========================== #

.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

# ========================= #
# Development
# ========================= #

## build: build the api binary
.PHONY: build
build:
	go build -o=/tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run: execute the build makefile command, then run the api binary
.PHONY: run
run:	build
			/tmp/bin/${BINARY_NAME} --port=${PORT}
