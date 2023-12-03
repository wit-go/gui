.PHONY: README.md log examples

all: README.md
	# reset
	@echo
	@echo "make examples     # will run all the Example demos and commands"
	@echo "make update       # full git update of all the dependencies"
	@echo
	@echo This Requires working IPv6
	@echo
	@sleep 1
ifeq (,$(wildcard go.mod))
	go mod init gui
	go mod tidy
endif
	make clean
	make plugins
	make examples-buttons

build-dep:
	apt install -f libgtk-3-dev

# should update every go dependancy (?)
update:
	git pull
	go get -v -t -u ./...

deb:
	cd debian && make
	dpkg-deb -c go-wit-gui*.deb
	-wit mirrors

examples:   \
	all \
	examples-helloworld \
	examples-buttons \
	examples-console-ui-helloworld \
	examples-textbox \
	examples-debug

examples-buttons:
	make -C examples/buttons

examples-console-ui-helloworld:
	make -C examples/console-ui-helloworld

# this is the most basic one. This syntax should always work
examples-helloworld:
	make -C examples/helloworld

examples-debug:
	-make -C examples/debug

examples-textbox:
	make -C examples/textbox

examples-helloconsole:
	make -C examples/plugin-consoleonly

# sync repo to the github backup
# git remote add github git@github.com:witorg/gui.git
# git remote add github2 git@github.com:wit-go/gui.git
github:
	git push origin master
	git push origin devel
	git push origin --tags
	git push github master
	git push github devel
	git push github --tags
	@echo
	@echo check https://github.com/wit-go/gui
	@echo

doc:
	godoc -v

goget:
	go get -v -t -u
	make -C toolkit/gocui goget
	make -C toolkit/andlabs goget

# GO111MODULE=on go install github.com/posener/goreadme/cmd/goreadme@latest (worked Oct 20 2022)
README.md: doc.go
	-goreadme -factories -types -functions -variabless > README-goreadme.md 

clean:
	rm -f toolkit/*.so
	cd debian && make clean

plugins: plugins-gocui plugins-andlabs

plugins-gocui:
	go build -C toolkit/gocui -v -buildmode=plugin -o ../gocui.so
	go build -C toolkit/nocui -v -buildmode=plugin -o ../nocui.so

plugins-andlabs:
	go build -C toolkit/andlabs -v -buildmode=plugin -o ../andlabs.so

objdump:
	objdump -t toolkit/andlabs.so |less

log:
	reset
	tail -f /tmp/witgui.* /tmp/guilogfile
