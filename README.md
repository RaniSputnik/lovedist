# lovedist

A tool for building Love 2d games for distribution. Following the instructions [here](https://love2d.org/wiki/Game_Distribution).

### Running the server

With [Docker](https://www.docker.com/) installed;

```
$ docker build -t lovedist .
$ docker run -p 8080:8080 lovedist
```

With [go](https://golang.org) installed;

```
$ go run *.go
```

### Tests

To run the tests use;

```
$ pushd ./handler; go test; popd;
```

A successful run will look something like the following;

```
~/go/src/github.com/RaniSputnik/lovedist/handler ~/go/src/github.com/RaniSputnik/lovedist
ok      github.com/RaniSputnik/lovedist/handler 1.220s
~/go/src/github.com/RaniSputnik/lovedist
```
