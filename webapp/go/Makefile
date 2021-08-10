GOPATH := ${PWD}
export GOPATH

install:
	dep ensure

build:
	go build -v isubata

vet:
	go vet ./src/isubata/...
