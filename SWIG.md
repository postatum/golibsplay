## Examples using SWIG

```sh
mkdir demos/swig/
```

### Node
```sh
$ mkdir demos/swig/node/
$ swig -outdir demos/swig/node/ -javascript -node mylib.i
$ cd demos/swig/node/
$ node
> (...)
```

### Java (SWIG)
```sh
$ swig -outdir demos/swig/java/ -java mylib.i
$ (...) # compile JAR
$ java -jar (...) # run JAR
```
