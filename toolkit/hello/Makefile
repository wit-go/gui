all: plugin
	ldd ../hello.so

build:
	GO111MODULE="off" go build

plugin:
	GO111MODULE="off" go build -buildmode=plugin -o ../hello.so
