.PHONY: README.md log examples

all: README.md
	# reset
	@echo
	@echo "make examples     # will run all the Example demos and commands"
	@echo "make update       # full git update of all the dependencies"
	@echo
	@echo This Requires working IPv6
	@echo
ifeq ($(GO111MODULE),)
	@echo
	@echo If you are compiling this here, you probably want to set GO111MODULE
	@echo
	@echo Setting GO111MODULE means that the version you are compiling has plugins
	@echo that get compiled against this current running version of the code
	@echo Otherwise, the GO language plugins can complain about being compiled against
	@echo mis-matched versions
	@echo
	@echo export GO111MODULE=off
	@echo
	sleep 3
endif
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

examples:   \
	all \
	examples-helloworld \
	examples-buttons \
	examples-console-ui-helloworld

# this is the most basic one. This syntax should always work
examples-helloworld:
	make -C examples/helloworld

examples-buttons:
	make -C examples/buttons

examples-console-ui-helloworld:
	make -C examples/console-ui-helloworld

git.wit.org:
	git push witgui master
	git push witgui devel
	git push witgui jcarr
	git push witgui --tags

# sync repo to the github backup
# git remote add github git@github.com:wit-go/gui.git
github: git.wit.org
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

clean:
	rm -f toolkit/*.so

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

submit-to-docs:
	GOPROXY=https://proxy.golang.org GO111MODULE=on go get go.wit.com/gui@v1.0.0
