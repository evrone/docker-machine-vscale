# Support go1.5 vendoring (let us avoid messing with GOPATH or using godep)
export GO15VENDOREXPERIMENT = 1

default: build

bin/docker-machine-driver-vscale: fetch
	go build -i -o ./bin/docker-machine-driver-vscale ./bin/

test: fetch
	go test ./...

fetch:
	go get -t -v ./...

build: clean bin/docker-machine-driver-vscale

clean:
	$(RM) -rf ./bin/docker*

install: bin/docker-machine-driver-vscale
	cp -f ./bin/docker-machine-driver-vscale $(GOPATH)/bin/ && \
	chmod +x $(GOPATH)/bin/docker-machine-driver-vscale

.PHONY: clean fetch test build bin/docker-machine-driver-vscale
