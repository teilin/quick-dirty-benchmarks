# Quick and dirty Benchmarks

## Int to string

```
go test -run=1000 -bench=.
goos: linux
goarch: amd64
pkg: devgeek.io/pkg/bench-int-string
cpu: Intel(R) Core(TM) Ultra 9 185H
BenchmarkIntStringSprintf-22    	25715224	        47.36 ns/op
BenchmarkIntStringItoa-22       	91297550	        13.74 ns/op
PASS
ok  	devgeek.io/pkg/bench-int-string	4.128s
```
