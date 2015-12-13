default: build

bin/docker-machine-driver-vscale:
	go build -i -o ./bin/docker-machine-driver-vscale ./bin/

build: clean bin/docker-machine-driver-vscale

clean:
	$(RM) -rf ./bin/docker*

install: bin/docker-machine-driver-vscale
	cp -f ./bin/docker-machine-driver-vscale $(GOPATH)/bin/ && \
	chmod +x $(GOPATH)/bin/docker-machine-driver-vscale

.PHONY: clean build
