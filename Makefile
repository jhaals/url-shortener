.PHONY: build push push-image

IMAGE_NAME ?= url-shortener
IMAGE_PREFIX ?= jhaals
IMAGE_TAG ?= 0.0.1

build:
	docker build -t $(IMAGE_PREFIX)/$(IMAGE_NAME):$(IMAGE_TAG) .

push: build push-image

push-image:
	docker push $(IMAGE_PREFIX)/$(IMAGE_NAME):$(IMAGE_TAG)
	docker push $(IMAGE_PREFIX)/$(IMAGE_NAME):latest
