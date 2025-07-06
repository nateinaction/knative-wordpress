REPO_NAME := knative-wordpress
REGISTRY ?= nateinaction
TAG_NAME ?= $(shell git rev-parse --short HEAD)$(shell git diff-index --quiet HEAD -- || echo "-dirty")

# Temporarily disable amd64 to improve local build times
# PLATFORMS=linux/amd64,linux/arm64
PLATFORMS=linux/arm64

build: build/php build/site

.PHONY: build/%
build/%:
	docker buildx build --platform $(PLATFORMS) -t ${REGISTRY}/${REPO_NAME}-$*:${TAG_NAME} images/$* --push
