all: plugin
	ldd ../nocui.so

plugin:
	GO111MODULE="off" go build -v -x -buildmode=plugin -o ../nocui.so
