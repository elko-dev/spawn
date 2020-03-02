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
MINOR_VERSION=5
PATCH_VERSION=1
CURRENT_VERSION=v$(MAJOR_VERSION).$(MINOR_VERSION).$(PATCH_VERSION)

default: make_start clean dependencies test build make_stop

make_start:
	@echo "make START"

clean: 
	$(GOCLEAN)
	rm -f wire_gen.go
	rm -rf mocks

dependencies:
	#global install
	go get github.com/google/wire/cmd/wire
	go get -u gopkg.in/src-d/go-git.v4
	go get github.com/golang/mock/mockgen@latest
	go install github.com/golang/mock/mockgen
	wire
	$(GOGET)
	mockgen -destination=applicationtype/applicationtype_mocks.go -package=applicationtype -source=applicationtype/applicationtypefactory.go -self_package=github.com/elko-dev/spawn/mocks
	mockgen -destination=web/web_mocks.go -package=web -source=web/factory.go -self_package=github.com/elko-dev/spawn/mocks
	mockgen -destination=applications/applications_mocks.go -package=applications -source=applications/applications.go -self_package=github.com/elko-dev/spawn/mocks
	mockgen -destination=nodejs/nodejs_mocks.go -package=nodejs -source=nodejs/factory.go -self_package=github.com/elko-dev/spawn/mocks
	mockgen -destination=herokuplatform/herokuplatform_mocks.go -package=herokuplatform -source=herokuplatform/factory.go -self_package=github.com/elko-dev/spawn/mocks
	mockgen -destination=git/git_mocks.go -package=git -source=git/factory.go -self_package=github.com/elko-dev/spawn/mocks
	mockgen -destination=platform/platform_mocks.go -package=platform -source=platform/factory.go -self_package=github.com/elko-dev/spawn/mocks

test: 
	$(GOTEST) -v ./...

build: spawn

spawn: $(GOFILES)	
	./scripts/build.sh $(APP_NAME) $(SRC_LOCATION) $(BIN_OUTPUT)

download:
	./scripts/download.sh $(CURRENT_VERSION)

publish-release:
	@go get github.com/aktau/github-release
	cd scripts && ./release.sh "v$(MAJOR_VERSION).$(MINOR_VERSION).$(PATCH_VERSION)" $(USER) $(APP_NAME) $(REPO)

make_stop:
	@echo "make STOP"
