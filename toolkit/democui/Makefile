all: plugin
	ldd ../democui.so

build:
	GO111MODULE="off" go build

plugin:
	GO111MODULE="off" go build -buildmode=plugin -o ../democui.so

objdump:
	objdump -t ../democui.so |less
