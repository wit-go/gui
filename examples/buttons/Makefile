#  with andlabs plugin loaded:
#     PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND
#  180006 jcarr     20   0 1918460  41688  31152 S   0.7   0.3   0:00.27 buttonplugin

#  with gocui plugin loaded:
#  180365 jcarr     20   0 1392668  24364  12596 S   2.0   0.2   0:00.09 buttonplugin
#

run: build
	./buttons

debug:
	./buttons --gui andlabs --gui-debug --log-debug

build-release:
	go get -v -u -x .
	go build
	./buttons

build:
	GO111MODULE="off" go get -v -x .
	GO111MODULE="off" go build

update:
	GO111MODULE="off" go get -v -u -x .

log:
	reset
	tail -f /tmp/witgui.* /tmp/guilogfile
