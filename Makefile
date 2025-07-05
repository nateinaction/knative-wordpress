REPO ?= nateinaction/knative-wordpress

PLATFORMS=linux/amd64,linux/arm64

build/%:
	docker buildx build --platform $(PLATFORMS) -t $(REPO)-$* images/$* --push
