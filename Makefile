# Commands
GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Configuration Values
SRC_LOCATION=.
BIN_OUTPUT=release
APP_NAME=spawn
REPO=elko-dev
MAJOR_VERSION=0
PATCH_VERSION=1
MINOR_VERSION=2

default: make_start clean dependencies test build make_stop

make_start:
	@echo "make START"

clean: 
	$(GOCLEAN)
	rm -f wire_gen.go

dependencies:
	#global install
	go get github.com/google/wire/cmd/wire
	go get -u gopkg.in/src-d/go-git.v4
	wire
	$(GOGET)

test: 
	$(GOTEST) -v ./...

build: spawn

spawn: $(GOFILES)	
	./scripts/build.sh $(APP_NAME) $(SRC_LOCATION) $(BIN_OUTPUT)

publish-release:
	@go get github.com/aktau/github-release
	cd scripts && ./release.sh "v$(MAJOR_VERSION).$(MINOR_VERSION).$(PATCH_VERSION)" $(USER) $(APP_NAME) $(REPO)

make_stop:
	@echo "make STOP"
