GIT_TAG?=$(shell git describe --always --tags)
BIN = rproxy
IMAGE_REPO = cosmincojocar
BUILDFLAGS := '-w -s'
CGO_ENABLED = 0
GO := GO111MODULE=on go
GO_NOMOD :=GO111MODULE=off go
GOPATH ?= $(shell $(GO) env GOPATH)
GOBIN ?= $(GOPATH)/bin
GOLINT ?= $(GOBIN)/golint
GOSEC ?= $(GOBIN)/gosec
GO_VERSION = 1.15

default:
	$(MAKE) build

test: build fmt lint sec
	 go test -v ./...

integration-test:
	tests/integration-tests.sh

fmt:
	@echo "FORMATTING"
	@FORMATTED=`$(GO) fmt ./...`
	@([[ ! -z "$(FORMATTED)" ]] && printf "Fixed unformatted files:\n$(FORMATTED)") || true

lint:
	@echo "LINTING"
	$(GO_NOMOD) get -u golang.org/x/lint/golint
	$(GOLINT) -set_exit_status ./...
	@echo "VETTING"
	$(GO) vet ./...

sec:
	@echo "SECURITY SCANNING"
	$(GO_NOMOD) get github.com/securego/gosec/cmd/gosec
	$(GOSEC) ./...

test-coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic

build:
	go build -o $(BIN) .

clean:
	rm -rf build vendor dist coverage.txt
	rm -f release image $(BIN)
	rm -rf tests/bin

build-linux:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 go build -ldflags $(BUILDFLAGS) -o $(BIN) .

image:
	@echo "Building the Docker image..."
	docker build -t $(IMAGE_REPO)/$(BIN):$(GIT_TAG) --build-arg GO_VERSION=$(GO_VERSION) .
	docker tag $(IMAGE_REPO)/$(BIN):$(GIT_TAG) $(IMAGE_REPO)/$(BIN):latest

image-push: image
	@echo "Pushing the Docker image..."
	docker push $(IMAGE_REPO)/$(BIN):$(GIT_TAG)
	docker push $(IMAGE_REPO)/$(BIN):latest

.PHONY: test integration-test build clean image image-push
