## Build stuff

```sh
$ make
```

## Demo examples

### Python
```sh
$ cp mylib.so demos/python/
$ cd demos/python/
$ python run.py
```

### Node
```sh
$ cd demos/node/
$ npm install
$ node index.js
```

### Browser (GopherJS)
```sh
$ make browser-example
$ open demos/browser/index.html
```

### Node (GopherJS)
```sh
$ make browser-example # edit myjslib.go first
$ node
> var lib = require('./myjslib.js')
> lib.PrintGoStr('hello')
```

### Java
```sh
cp mylib.so demos/java/
cd demos/java/
mvn install
java -cp target/classes:target/dependency/* javatogo.JavaToGo
```

### CLI
```sh
go get github.com/codegangsta/cli
cd mycli
go install
cd ../demos
mycli jsonfirst ./testdata.json
```
