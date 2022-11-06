.PHONY: README.md

all: README.md
	@echo
	@echo "make examples     # will run all the examples"
	@echo "make update       # full git update"
	@echo
	make -C cmds/helloworld

# should update every go dependancy (?)
update:
	git pull
	go get -v -t -u ./...

examples: examples-consolemouse examples-helloworld examples-gui-demo

examples-consolemouse:
	make -C cmds/consolemouse

examples-helloworld:
	make -C cmds/helloworld

examples-gui-demo:
	make -C cmds/gui-demo

# sync repo to the github backup
github:
	git push origin master
	git push github master

doc:
	GO111MODULE="off" godoc -v


# GO111MODULE=on go install github.com/posener/goreadme/cmd/goreadme@latest (worked Oct 20 2022)
README.md: doc.go
	goreadme -factories -types -functions -variabless > README.md 

plugins:
	GO111MODULE="off" go build -buildmode=plugin -o toolkit/gocli.so toolkit/gocli/greeter.go
