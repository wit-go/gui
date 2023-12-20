all:
	# You must use the current protoc-gen-go
	# protoc --version 3.6++ does not mean that protoc will generate version3 .go files
	#
	# apt remove golang-goprotobuf-dev
	# apt install protobuf-compiler
	#
	# Then:
	# go get -u github.com/golang/protobuf/protoc-gen-go
	# cd ~/go/src/github.com/golang/protobuf/protoc-gen-go
	# go install
	#
	# Then:
	protoc --version
	make widget.pb.go

clean:
	rm -f *.pb.go

widget.pb.go: widget.proto
	protoc --go_out=. widget.proto

compile:
	protoc --go_out=. *.proto

deps:
	apt install golang-goprotobuf-dev
	apt install protobuf-compiler

push:
	git pull
	git add --all
	git commit -a -s
	git push
