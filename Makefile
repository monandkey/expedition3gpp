# Makefile
DOCKER_ENV                ?= DOCKER_BUILDKIT=1
DOCKER_TAG                ?= 0.2
DOCKER_REGISTRY           ?= docker.io
DOCKER_REPOSITORY         ?= monandkey/expedition3gpp
DOCKER_BUILD_ARGS         ?= --rm

EXPEDITION3GPP_IMAGE_NAME ?= ${DOCKER_REGISTRY}/${DOCKER_REPOSITORY}:${DOCKER_TAG}


build-all: build-ueransim

.PHONY: build-ueransim
build-ueransim: 
	${DOCKER_ENV} docker build ${DOCKER_BUILD_ARGS} \
		--tag ${EXPEDITION3GPP_IMAGE_NAME} \
		--file ./Dockerfile \
		--no-cache \
		./

# Sample Command
# DOCKER_BUILDKIT=1 docker build --rm \
#   --tag docker.io/monandkey/expedition3gpp:0.2 \
#   --file ./Dockerfile \
#   --no-cache \
#   ./