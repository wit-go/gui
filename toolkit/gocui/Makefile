all: plugin
	ldd ../gocui.so

goget:
	go get -v -t -u

plugin:
	go build -v -x -buildmode=plugin -o ../gocui.so

objdump:
	objdump -t ../gocui.so |less

log:
	reset
	tail -f /tmp/witgui.* /tmp/guilogfile
