# should update every go dependancy (?)
update:
	git pull
	GO111MODULE="off" go get -v -t -u ./...
