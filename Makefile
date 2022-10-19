all:
	@echo
	@echo "make examples     # will run all the examples"
	@echo "make update       # full git update"
	@echo

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

doc:
	GO111MODULE="off" godoc -v
