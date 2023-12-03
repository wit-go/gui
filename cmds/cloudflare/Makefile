run: build
	./cloudflare

build-release:
	go get -v -u -x .
	go build
	./cloudflare

build:
	GO111MODULE="off" go get -v -x .
	GO111MODULE="off" go build

update:
	GO111MODULE="off" go get -v -u -x .

log:
	reset
	tail -f /tmp/witgui.* /tmp/guilogfile
