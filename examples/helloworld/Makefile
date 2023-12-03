run: build
	./helloworld

build-release:
	go get -v -u -x .
	go build
	./helloworld

build:
	GO111MODULE="off" go get -v -x .
	GO111MODULE="off" go build

update:
	GO111MODULE="off" go get -v -u -x .
