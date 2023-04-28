all: plugin

plugin:
	GO111MODULE="off" go build -v -x -buildmode=plugin -o ../andlabs.so

goget:
	GO111MODULE="off" go get -v -t -u
