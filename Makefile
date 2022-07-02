.PHONY: all build

BUILDTIME ?= $(shell date +%Y-%m-%d_%I:%M:%S)
GITCOMMIT ?= $(shell git rev-parse -q HEAD)
ifeq ($(CI_PIPELINE_ID),)
	BUILDNUMER := private
else
	BUILDNUMER := $(CI_PIPELINE_ID)
endif
VERSION ?= $(shell git describe --tags --always --dirty)

LDFLAGS = -extldflags \
		  -static \
		  -X "main.Version=$(VERSION)" \
		  -X "main.BuildTime=$(BUILDTIME)" \
		  -X "main.GitCommit=$(GITCOMMIT)" \
		  -X "main.BuildNumber=$(BUILDNUMER)"

all: build

build:
	go build cmd/main.go;
	mv main bin/gateway

