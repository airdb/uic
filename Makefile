SHELL = /bin/bash

VERSION:=$(shell git describe --dirty --always)
#VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse HEAD)
REPO := github.com/airdb/uic

LDFLAGS=-ldflags
LDFLAGS += "-X=github.com/airdb/sailor/version.Repo=$(REPO) \
            -X=github.com/airdb/sailor/version.Version=$(VERSION) \
            -X=github.com/airdb/sailor/version.Build=$(BUILD) \
            -X=github.com/airdb/sailor/version.BuildTime=$(shell date +%s)"

# SLSENV=SERVERLESS_PLATFORM_VENDOR=tencent GLOBAL_ACCELERATOR_NA=true
SLSENV=SERVERLESS_PLATFORM_VENDOR=tencent
default: build deploy

build:
	GOOS=linux go build $(LDFLAGS) -o main main.go

swag:
	swag init --generalInfo internal/app/adapter/controller.go --output ./docs

dev:
	env=dev go run  $(LDFLAGS) main.go

wire:
	wire gen internal/app/wire.go

deploy: swag
	${SLSENV} sls deploy --stage test
	@echo checkout all scf apps, https://serverless.cloud.tencent.com/

release: swag
	${SLSENV} sls deploy --stage release
	@echo checkout all scf apps, https://serverless.cloud.tencent.com/

log:
	${SLSENV} sls logs --tail --stage test

logrelease:
	${SLSENV} sls logs --tail --stage release
