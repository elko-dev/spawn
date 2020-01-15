GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

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

make_stop:
	@echo "make STOP"
