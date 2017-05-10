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