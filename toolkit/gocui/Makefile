all: plugin
	ldd ../gocli.so

build:
	GO111MODULE="off" go build

plugin:
	GO111MODULE="off" go build -buildmode=plugin -o ../gocli.so
