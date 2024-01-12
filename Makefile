all:
	@echo
	@echo This is the core gui package 'go.wit.com/gui/gui'
	@echo
	@echo It creates a binary tree of widgets
	@echo The widgets are things like Windows, Buttons, Labels, etc
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
endif

redomod:
	rm -f go.*
	GO111MODULE= go mod init
	GO111MODULE= go mod tidy

# should update every go dependancy (?)
update:
	git pull
	go get -v -t -u ./...

# sync repo to the github backup
# git remote add github git@github.com:wit-go/gui.git
github:
	git push origin master
	git push origin devel
	git push origin --tags
	git push github master
	git push github devel
	git push github --tags
	@echo
	@echo check https://git.wit.org/gui/gui
	@echo check https://github.com/wit-go/gui
	@echo

doc:
	godoc -v

submit-to-docs:
	GOPROXY=https://proxy.golang.org GO111MODULE=on go get go.wit.com/gui@v1.0.0
