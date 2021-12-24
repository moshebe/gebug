WORKROOT := $(shell pwd)
OUTDIR   := $(WORKROOT)/output

export PATH        := $(shell go env GOPATH)/bin:$(PATH)
export GO111MODULE := on

GOFLAGS      := -race
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
	go build -ldflags "-X main.version=$(GEBUG_VERSION) -X main.commit=$(GIT_COMMIT)"

test: test-case vet-case
test-case:
	go test -cover ./...
vet-case:
	go vet ./...

coverage:
	echo -n > coverage.txt
	for pkg in $(PKGS) ; do go test -coverprofile=profile.out -covermode=atomic $${pkg} && cat profile.out >> coverage.txt; done

check:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...

clean:
	rm -rf $(OUTDIR)
	rm -rf $(WORKROOT)/gebug
	rm -rf $(WORKROOT)/.gebug
	rm -rf $(GOPATH)/pkg/linux_amd64

.PHONY: all compile test package clean build