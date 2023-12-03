run: build
	./console-ui-helloworld
	reset
	ldd ./console-ui-helloworld

build-release:
	go get -v -u -x .
	go build

build:
	GO111MODULE="off" go get -v -x .
	GO111MODULE="off" go build

update:
	GO111MODULE="off" go get -v -u -x .
