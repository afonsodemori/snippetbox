include .env

# development

run:
	go run ./cmd/web

certs:
	@mkdir -p tls
	@cd tls && go run $(shell go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost

# docker

COMMIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
IMAGE_TAG := $(shell date +%Y%m%d-%H%M%S)

docker-build:
	@echo "Building Docker image $(IMAGE_NAME) with GO_VERSION=$(GO_VERSION)"
	docker buildx build . \
		--platform=linux/amd64,linux/arm64 \
		--file=Dockerfile \
		--build-arg GO_VERSION=$(GO_VERSION) \
		--tag=$(IMAGE_NAME):$(IMAGE_TAG) \
		--tag=$(IMAGE_NAME):latest

docker-push: docker-build
	echo "Pushing tag: $(IMAGE_TAG)..."
	docker push $(IMAGE_NAME):$(IMAGE_TAG)
	echo "Pushing tag: latest..."
	docker push $(IMAGE_NAME):latest
