# gorand [![Godoc](https://godoc.org/github.com/yuansip/gorand?status.svg)](https://godoc.org/github.com/yuansip/gorand)

gorand is a Go library for selecting n elements from element slice randomly at one time according to specific element probability distribution

## Installation

Standard `go get`:

```
$ go get github.com/yuansip/gorand
```

## Benchmarks

###### Run on Mac mini (i5, 8G, 2014) go version go1.11.4 darwin/amd64

```go
go test -bench=. -benchmem ./...
goos: darwin
goarch: amd64
pkg: github.com/yuansip/gorand
BenchmarkRandomSelect100-4         50000             32161 ns/op            7312 B/op          7 allocs/op
BenchmarkRandomSelect1000-4         5000            339758 ns/op           21904 B/op          7 allocs/op
BenchmarkRandomSelect10000-4         300           4355937 ns/op          169360 B/op          7 allocs/op
PASS
ok      github.com/yuansip/gorand       6.986s

```

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.
