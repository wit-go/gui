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

merge-devel:
	git checkout master
	git pull origin master
	git merge devel
	git push origin master
	git checkout devel
