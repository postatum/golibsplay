## Build stuff

### C shared library

```sh
$ go build -o mylib.so -buildmode=c-shared mylib.go
```

### JS library
```sh
$ go get -u github.com/gopherjs/gopherjs
$ gopherjs build myjslib.go
```

## Demo examples

### Node (GopherJS)
```sh
$ gopherjs build myjslib.go
$ mv myjslib.js* demos/node-gopherjs/
$ cd demos/node-gopherjs/
$ node
> var lib = require('./myjslib.js')
> lib.PrintGoStr('hello')
```

### Node (Shared library)
```sh
$ cp mylib.so mylib.so.dylib # OSX hack for node-ffi
$ cd demos/node/
$ npm install
$ node index.js
```

### Node (SWIG)
```sh
$ swig -outdir demos/node-swig/ -javascript mylib.i
$ cd demos/node-swig/
$ node
> (...)
```

### Browser (GopherJS)
```sh
$ gopherjs build myjslib.go # edit myjslib.go first
$ mv myjslib.js* demos/browser/
$ open demos/browser/index.html
```

### Java (SWIG)
```sh
$ swig -outdir demos/java/ -java mylib.i
$ (...) # compile JAR
$ java -jar (...) # run JAR
```
