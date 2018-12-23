# Public IPFS Gateway Checker in Golang

This's a Golang migration from:

* [public-gateway-checker](https://github.com/ipfs/public-gateway-checker) A online webpage of JS version
* [ipfg](https://github.com/JayBrown/Tools/tree/master/ipfg)  A shell cli version

### Install

```bash
go get github.com/gingerhot/ipfg
```

### Usage

After installation, there will be a command `ipfg` available for you.

```bash
ipfg        # to get one available gateway url for you
ipfg all    # get all available gateway urls
ipfg help   # show help message
```
And you can import `github.com/gingerhot/utils` in your code, just as what I do in the [main.go](./main.go).
There're two methods for you: `Get()` and `ActiveList()`. More detail: [godoc](https://godoc.org/github.com/gingerhot/ipfg)

