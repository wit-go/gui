all: plugin
	ldd ../gocui.so

build:
	GO111MODULE="off" go build

plugin:
	GO111MODULE="off" go build -buildmode=plugin -o ../gocui.so

objdump:
	objdump -t ../gocui.so |less
