# gorand [![Godoc](https://godoc.org/github.com/yuansip/gorand?status.svg)](https://godoc.org/github.com/yuansip/gorand)

gorand is a Go library for selecting n elements from element slice randomly at one time. The slice can have a specific element probability distribution

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
BenchmarkRandomSelect1_100-4                       50000             31984 ns/op            7264 B/op          5 allocs/op
BenchmarkRandomSelect1_1000-4                       5000            339030 ns/op           21856 B/op          5 allocs/op
BenchmarkRandomSelect1_10000-4                       300           4351709 ns/op          169312 B/op          5 allocs/op
BenchmarkRandomSelect50_100-4                       2000            575657 ns/op          270688 B/op         54 allocs/op
BenchmarkRandomSelect500_1000-4                      200           6707793 ns/op         2704480 B/op        504 allocs/op
BenchmarkRandomSelect5000_10000-4                     10         256038988 ns/op        27043945 B/op       5004 allocs/op
BenchmarkRandomSelect99_100-4                       1000           1115691 ns/op          534112 B/op        103 allocs/op
BenchmarkRandomSelect999_1000-4                      100          12600892 ns/op         5387104 B/op       1003 allocs/op
BenchmarkRandomSelect9999_10000-4                      5         252467309 ns/op        53918560 B/op      10003 allocs/op
BenchmarkRandomSelectEvenly1_100-4                200000             10917 ns/op            5376 B/op          1 allocs/op
BenchmarkRandomSelectEvenly1_1000-4               200000             10926 ns/op            5376 B/op          1 allocs/op
BenchmarkRandomSelectEvenly1_10000-4              200000             11020 ns/op            5376 B/op          1 allocs/op
BenchmarkRandomSelectEvenly50_100-4                 3000            551184 ns/op          268800 B/op         50 allocs/op
BenchmarkRandomSelectEvenly500_1000-4                300           5474139 ns/op         2688000 B/op        500 allocs/op
BenchmarkRandomSelectEvenly5000_10000-4               20          55147806 ns/op        26880000 B/op       5000 allocs/op
BenchmarkRandomSelectEvenly99_100-4                 2000           1087750 ns/op          532224 B/op         99 allocs/op
BenchmarkRandomSelectEvenly999_1000-4                100          10962265 ns/op         5370624 B/op        999 allocs/op
BenchmarkRandomSelectEvenly9999_10000-4               10         110186931 ns/op        53754624 B/op       9999 allocs/op
PASS
ok      github.com/yuansip/gorand       36.063s

```

## Examples

```go
elements := []int{1,2,3,4,5}
selection := gorand.SelectIntEvenly(elements)
selections := gorand.SelectNIntEvenly(elements, 2)

type Selectable struct {
	Value       int
	Probability float64
}

func (s Selectable) GetValue() interface{} {
	return s.Value
}

func (s Selectable) GetProbability() float64 {
	return s.Probability
}
elements2 := []ProbabilityElement{
	Selectable{
	    Value:1,
	    Probability:0.2,
	},
	Selectable{
	    Value:2,
	    Probability:0.1,
	},
	Selectable{
	    Value:3,
	    Probability:0.3,
	},
	Selectable{
	    Value:4,
	    Probability:0.3,
	},
	Selectable{
	    Value:5,
	    Probability:0.1,
	},
}

sel := gorand.RandomSelect(elements2)
sels := gorand.RandomSelectN(elements2, 2)
```

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.
