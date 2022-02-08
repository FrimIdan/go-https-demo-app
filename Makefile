e_Y=\033[1;33m
C_C=\033[0;36m
C_M=\033[0;35m
C_R=\033[0;41m
C_N=\033[0m
SHELL=/bin/bash

# Project variables
BINARY_NAME ?= https-test
DOCKER_REGISTRY ?= idanfrim
VERSION ?= $(shell git rev-parse HEAD)
DOCKER_IMAGE ?= $(DOCKER_REGISTRY)/$(BINARY_NAME)
DOCKER_TAG ?= ${VERSION}

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help api https-test docker docker-push clean clean-https-test keys

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

.PHONY: api
api: ## Generating API code
	@(echo "Generating API code ..." )
	@./generate.sh

.PHONY: https-test
https-test: ## Build HTTPs Test App
	@(echo "Building HTTPs Test App ..." )
	@(go build -o bin/https-test cmd/https-test/main.go && ls -l bin/)

.PHONY: docker
docker: ## Build Docker image
	@(echo "Building docker image ..." )
	@(docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG} .)

.PHONY: docker-push
docker-push: docker
	@echo "Publishing Docker image ..."
	@(docker push ${DOCKER_IMAGE}:${DOCKER_TAG})

.PHONY: clean
clean:
	@(rm -rf bin/https-test ; echo "Cleanup done" )

.PHONY: keys
keys:
	# The CName used here is specific to the service specified in deploy.yaml.
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout $(KEY) -out $(CERT) -subj "/CN=server.default/O=server"
