.PHONY: help build
.DEFAULT_GOAL := help

BUILD_DIR := "build"

run: build  ## Run the application
	./${BUILD_DIR}/goflix

build:  ## Build the application in the build/ directory
	# mkdir -p ${BUILD_DIR}	
	go build -o ${BUILD_DIR}/

watch:  ## Automatic build & run on file update
	ls *.go | entr -c -r make run

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
