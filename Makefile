.PHONY: README.md log

all: README.md
	reset
	@echo
	@echo "make examples     # will run all the Example demos and commands"
	@echo "make update       # full git update of all the dependencies"
	@echo
	make clean
	make plugins

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
	cmds-helloworld \
	cmds-buttonplugin \
	cmds-console-ui-helloworld \
	cmds-textbox \
	cmds-debug

cmds-buttonplugin:
	make -C cmds/buttonplugin

cmds-console-ui-helloworld:
	make -C cmds/console-ui-helloworld

# this is the most basic one. This syntax should always work
cmds-helloworld:
	make -C cmds/helloworld

cmds-debug:
	-make -C cmds/debug

cmds-textbox:
	make -C cmds/textbox

cmds-helloconsole:
	make -C cmds/plugin-consoleonly

# sync repo to the github backup
# git remote add github git@github.com:witorg/gui.git
# git remote add github2 git@github.com:wit-go/gui.git
github:
	git push origin master
	git push github master
	git push github devel
	git push github --tags
	@echo
	@echo check https://github.com/wit-go/gui
	@echo

doc:
	GO111MODULE="off" godoc -v


# GO111MODULE=on go install github.com/posener/goreadme/cmd/goreadme@latest (worked Oct 20 2022)
README.md: doc.go
	-goreadme -factories -types -functions -variabless > README-goreadme.md 

clean:
	rm -f toolkit/*.so
	cd debian && make clean

plugins: plugins-gocui plugins-andlabs

plugins-gocui:
	make -C  toolkit/gocui

plugins-andlabs:
	cd toolkit/andlabs/ && GO111MODULE="off" go build -buildmode=plugin -o ../andlabs.so
	# make -C  toolkit/andlabs

objdump:
	objdump -t toolkit/andlabs.so |less

log:
	reset
	tail -f /tmp/witgui.* /tmp/guilogfile
