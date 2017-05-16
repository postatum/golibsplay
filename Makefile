all:
	go build -o mylib.so -buildmode=c-shared mylib.go
	cp mylib.so mylib.so.dylib # OSX hack for node-ffi

browser-example:
	go get -u github.com/gopherjs/gopherjs
	gopherjs build myjslib.go
	mv myjslib.js* demos/browser/

node-gopherjs-example:
	go get -u github.com/gopherjs/gopherjs
	mkdir demos/node-gopherjs/
	gopherjs build myjslib.go
	mv myjslib.js* demos/node-gopherjs/

clean:
	rm -f mylib.h mylib.so