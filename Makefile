# simple sortcut to push all git changes
push:
	git pull
	git add --all
	-git commit -a -s
	git push

# should update every go dependancy (?)
update:
	git pull
	go get -v -t -u ./...
