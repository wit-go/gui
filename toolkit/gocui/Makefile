all: plugin
	ldd ../gocui.so

goget:
	GO111MODULE="off" go get -v -t -u

build:
	GO111MODULE="off" go build

plugin:
	GO111MODULE="off" go build -buildmode=plugin -o ../gocui.so

objdump:
	objdump -t ../gocui.so |less

log:
	reset
	tail -f /tmp/witgui.* /tmp/guilogfile
