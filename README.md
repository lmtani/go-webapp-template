# Go webapp template

Go webapp with React and embebbed UI files.

I'm using [sausheong/invadersapp](https://github.com/sausheong/invadersapp) as a starter point.

### Build

```sh
$ node --version
v15.10.0

$ go version
go version go1.16.5 linux/amd64

$ cd ui && npm run build && cd -
$ go build -ldflags="-s -w"
```
