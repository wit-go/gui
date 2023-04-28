all: plugin
	ldd ../gocui.so

goget:
	GO111MODULE="off" go get -v -t -u

plugin:
	GO111MODULE="off" go build -v -x -buildmode=plugin -o ../gocui.so

objdump:
	objdump -t ../gocui.so |less

log:
	reset
	tail -f /tmp/witgui.* /tmp/guilogfile
