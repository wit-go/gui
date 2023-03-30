run: build
	# GOTRACEBACK=all ./textbox --gui-debug
	GOTRACEBACK=all ./textbox --gui-debug >/tmp/witgui.log.stderr 2>&1

build-release:
	go get -v -u -x .
	go build

build:
	GO111MODULE="off" go get -v -x .
	GO111MODULE="off" GOTRACEBACK=all go build

update:
	GO111MODULE="off" go get -v -u -x .
