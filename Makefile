# Makefile params
SHELL             = /bin/bash -o pipefail

WORKROOT := $(shell pwd)
OUTDIR   := $(WORKROOT)/output

GOARCH            = amd64
GOOS              ?= $(shell go env GOOS)
GOPATH            ?= $(shell go env GOPATH)
COMMIT            ?= $(shell git rev-parse HEAD)
BRANCH            ?= $(shell git rev-parse --abbrev-ref HEAD)
BUILD_DATE        ?= $(shell date -u +"%Y-%m-%d")
VERSION           ?= $(shell git describe --long --tags)
BASENAME          ?= gebug
BINARY            ?= ${BASENAME}
VERSION_PKG   	  = github.com/moshebe/gebug/version
LDFLAGS	          = -ldflags \
					"-X ${VERSION_PKG}.Version=${VERSION} -X ${VERSION_PKG}.Revision=${COMMIT} -X ${VERSION_PKG}.Branch=${BRANCH}"

STATICCHECK  := staticcheck

ARCH := $(shell getconf LONG_BIT)
ifeq ($(ARCH),64)
	GOTEST += -race
endif

GEBUG_VERSION ?= $(shell cat VERSION)
GIT_COMMIT ?= $(shell git rev-parse HEAD)

PKGS := $(shell go list ./...)

all: compile package

compile: test build
build:
	go build ${LDFLAGS}

.PHONY: buildall
	$(MAKE) build GOOS=windows BINARY=${BINARY}-windows-${GOARCH}.exe
	$(MAKE) build GOOS=linux BINARY=${BINARY}-linux-${GOARCH}
	$(MAKE) build GOOS=darwin BINARY=${BINARY}-darwin-${GOARCH}

test: test-case vet-case
test-case:
	go test -cover ./...
vet-case:
	go vet ./...

coverage:
	echo -n > coverage.txt
	for pkg in $(PKGS) ; do go test -coverprofile=profile.out -covermode=atomic $${pkg} && cat profile.out >> coverage.txt; done

package:
	mkdir -p $(OUTDIR)/bin
	mv bfe  $(OUTDIR)/bin
	cp -r conf $(OUTDIR)

check:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...

clean:
	rm -rf $(OUTDIR)
	rm -rf $(WORKROOT)/gebug
	rm -rf $(WORKROOT)/.gebug
	rm -rf $(GOPATH)/pkg/linux_amd64

.PHONY: all compile test package clean build buildall