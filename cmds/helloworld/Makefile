run: build
	./helloworld

build-release:
	go get -v -u -x .
	go build

build:
	GO111MODULE="off" go get -v -x .
	GO111MODULE="off" go build
	./helloworld

update:
	GO111MODULE="off" go get -v -u -x .
