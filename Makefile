# Makefile
DOCKER_ENV                ?= DOCKER_BUILDKIT=1
DOCKER_TAG                ?= 0.2
DOCKER_REGISTRY           ?= docker.io
DOCKER_REPOSITORY         ?= monandkey/expedition3gpp
DOCKER_BUILD_ARGS         ?= --rm

EXPEDITION3GPP_IMAGE_NAME ?= ${DOCKER_REGISTRY}/${DOCKER_REPOSITORY}:${DOCKER_TAG}


build-all: build-expedition3gpp

.PHONY: build-expedition3gpp
build-expedition3gpp: 
	${DOCKER_ENV} docker build ${DOCKER_BUILD_ARGS} \
		--tag ${EXPEDITION3GPP_IMAGE_NAME} \
		--file ./build/Dockerfile \
		--no-cache \
		./

# Sample Command
# DOCKER_BUILDKIT=1 docker build --rm \
#   --tag docker.io/monandkey/expedition3gpp:0.2 \
#   --file ./build/Dockerfile \
#   --no-cache \
#   ./