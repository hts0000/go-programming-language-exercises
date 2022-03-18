cmdline run ```go test -cpu=4 -bench .``` to start benchmak test

**test result**

goos: windows
goarch: amd64
pkg: execrises-2.3
cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
BenchmarkPopCount1-4    1000000000               0.2553 ns/op
BenchmarkPopCount2-4    354241450                3.337 ns/op
PASS
ok      execrises-2.3   2.470s