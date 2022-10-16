run: build
	./helloworld

build:
	go get -v -u -x .
	go build

build-master:
	GO111MODULE="off" go get -v -x .
	GO111MODULE="off" go build
	./helloworld

update:
	GO111MODULE="off" go get -v -u -x .
