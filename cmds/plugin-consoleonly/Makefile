run: build
	./helloconsole

build-release:
	go get -v -u -x .
	go install -v -o helloconsole
	helloconsole

build:
	GO111MODULE="off" go get -v -x .
	GO111MODULE="off" go build -v -o helloconsole

update:
	GO111MODULE="off" go get -v -u -x .
