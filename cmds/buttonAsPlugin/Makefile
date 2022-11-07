run: build
	./buttonAsPlugin >/tmp/buttonAsPlugin.log 2>&1

build-release:
	go get -v -u -x .
	go build
	./buttonAsPlugin

build:
	GO111MODULE="off" go get -v -x .
	GO111MODULE="off" go build

update:
	GO111MODULE="off" go get -v -u -x .
