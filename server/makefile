GO_VERSION ?= 1.16
IMAGE_NAME ?= ghcr.io/reliablyhq/walkthrough-slo-http-server:dev

.PHONY: build
build: install test
	GOOS=linux go build -o bin/server main.go

.PHONY: install
install:
	go mod vendor

.PHONY: test
test:
	go test ./...

.PHONY: build-image
build-image:
	docker build -t ${IMAGE_NAME} --build-arg GO_VERSION=${GO_VERSION} .

.PHONY: build-and-push-image
build-and-push-image: build-image
	docker push ${IMAGE_NAME}