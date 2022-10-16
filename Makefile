all:
	@echo
	@echo "make examples     # will run all the examples"
	@echo "make update       # full git update"
	@echo

# should update every go dependancy (?)
update:
	git pull
	go get -v -t -u ./...

examples:
	make -C cmds/helloworld
	make -C cmds/gui-example
	make -C cmds/gui-demo

doc:
	godoc -v
