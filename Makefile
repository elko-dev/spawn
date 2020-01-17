GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

APP_NAME=spawn
MAJOR_VERSION=0
MINOR_VERSION=1

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

build: 
	$(GOBUILD) 

publish-release:
	@go get github.com/aktau/github-release
	cd release && ./release.sh "v$(MAJOR_VERSION).$(MINOR_VERSION).$(VERSION)" $(APP_NAME) $(USER)

make_stop:
	@echo "make STOP"
