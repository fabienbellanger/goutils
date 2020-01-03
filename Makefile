# Go parameters
GOCMD=go
GOINSTALL=$(GOCMD) install
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOTOOL=$(GOCMD) tool

all: test

install:
	$(GOINSTALL) ./...

update:
	$(GOGET) -u && $(GOMOD) tidy

update-all:
	$(GOGET) -u all && $(GOMOD) tidy

test: 
	$(GOTEST) -cover ./...

test-cover-count: 
	$(GOTEST) -covermode=count -coverprofile=cover-count.out ./...

test-cover-atomic: 
	$(GOTEST) -covermode=atomic -coverprofile=cover-atomic.out ./...

cover-count:
	$(GOTOOL) cover -func=cover-count.out

cover-atomic:
	$(GOTOOL) cover -func=cover-atomic.out

html-cover-count:
	$(GOTOOL) cover -html=cover-count.out

html-cover-atomic:
	$(GOTOOL) cover -html=cover-atomic.out

run-cover-count: test-cover-count cover-count
run-cover-atomic: test-cover-atomic cover-atomic
view-cover-count: test-cover-count html-cover-count
view-cover-atomic: test-cover-atomic html-cover-atomic

bench: 
	$(GOTEST) -bench=. ./...

clean: 
	$(GOCLEAN)
