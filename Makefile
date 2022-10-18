all:
	@echo
	@echo "make examples     # will run all the examples"
	@echo "make update       # full git update"
	@echo

# should update every go dependancy (?)
update:
	git pull
	go get -v -t -u ./...

examples: examples-gui-demo examples-helloworld

examples-helloworld:
	make -C cmds/helloworld

examples-gui-demo:
	make -C cmds/gui-demo

doc:
	godoc -v
