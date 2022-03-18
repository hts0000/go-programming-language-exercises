go test -cpu 4 -bench . -benchmem
goos: windows
goarch: amd64
pkg: execrises-2.4
cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
BenchmarkPopCount1-4    1000000000               0.2525 ns/op          0 B/op          0 allocs/op
BenchmarkPopCount2-4    54545702                21.88 ns/op            0 B/op          0 allocs/op
BenchmarkPopCount3-4    10117693               111.6 ns/op             0 B/op          0 allocs/op
PASS
ok      execrises-2.4   3.379s