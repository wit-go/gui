.PHONY: README.md

all: README.md
	@echo
	@echo "make cmds     # will run all the Example demos and commands"
	@echo "make update       # full git update of all the dependencies"
	@echo
	#make -C cmds/helloworld
	make plugins

# should update every go dependancy (?)
update:
	git pull
	go get -v -t -u ./...

cmds: cmds-buttonplugin cmds-console-ui-helloworld cmds-debug cmds-helloworld cmds-textbox

cmds-buttonplugin:
	make -C cmds/buttonplugin

cmds-console-ui-helloworld:
	make -C cmds/console-ui-helloworld

cmds-helloworld:
	make -C cmds/helloworld

cmds-debug:
	make -C cmds/debug

cmds-textbox:
	make -C cmds/textbox

# sync repo to the github backup
# git remote add github git@github.com:witorg/gui.git
github:
	git push origin master
	git push github master
	git push github devel
	git push github jcarr
	@echo
	@echo check https://github.com/witorg/gui
	@echo

doc:
	GO111MODULE="off" godoc -v


# GO111MODULE=on go install github.com/posener/goreadme/cmd/goreadme@latest (worked Oct 20 2022)
README.md: doc.go
	goreadme -factories -types -functions -variabless > README-goreadme.md 

clean:
	rm -f toolkit/*.so

plugins: plugins-gocui plugins-andlabs

plugins-gocui:
	make -C  toolkit/gocui

plugins-andlabs:
	cd toolkit/andlabs/ && GO111MODULE="off" go build -buildmode=plugin -o ../andlabs.so
	# make -C  toolkit/andlabs

objdump:
	objdump -t toolkit/andlabs.so |less
