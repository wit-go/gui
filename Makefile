.PHONY: README.md

all: README.md
	@echo
	@echo "make cmds     # will run all the cmds"
	@echo "make update       # full git update"
	@echo
	make -C cmds/helloworld

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
github:
	git push origin master
	git push github master

doc:
	GO111MODULE="off" godoc -v


# GO111MODULE=on go install github.com/posener/goreadme/cmd/goreadme@latest (worked Oct 20 2022)
README.md: doc.go
	goreadme -factories -types -functions -variabless > README-goreadme.md 

clean:
	rm -f toolkit/*.so

plugins:
	# GO111MODULE="off" go build -buildmode=plugin -o toolkit/test.so toolkit/gocui/*.go
	make -C  toolkit/gocui
