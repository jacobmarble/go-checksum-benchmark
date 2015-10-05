I needed a fast, not necessarily crypto-secure, checksum algorithm. Here are my results:

```sh
$ go test -bench . -benchtime 10s -benchmem
testing: warning: no tests to run
PASS
BenchmarkMD4-24       	      50	 348684216 ns/op	      72 B/op	       2 allocs/op
BenchmarkMD5-24       	     100	 149806087 ns/op	      36 B/op	       1 allocs/op
BenchmarkSHA1-24      	     100	 201456653 ns/op	      52 B/op	       1 allocs/op
BenchmarkSHA224-24    	      30	 505730641 ns/op	      36 B/op	       1 allocs/op
BenchmarkSHA256-24    	      30	 504088949 ns/op	      36 B/op	       1 allocs/op
BenchmarkSHA384-24    	      50	 307163834 ns/op	      52 B/op	       1 allocs/op
BenchmarkSHA512-24    	      50	 308922151 ns/op	      68 B/op	       1 allocs/op
BenchmarkSHA3_224-24  	      30	 387445669 ns/op	     526 B/op	       3 allocs/op
BenchmarkSHA3_256-24  	      30	 407159381 ns/op	     526 B/op	       3 allocs/op
BenchmarkSHA3_384-24  	      30	 539065391 ns/op	     558 B/op	       3 allocs/op
BenchmarkSHA3_512-24  	      20	 788527154 ns/op	     598 B/op	       3 allocs/op
BenchmarkSHA512_224-24	      50	 316293949 ns/op	      36 B/op	       1 allocs/op
BenchmarkSHA512_256-24	      50	 311169666 ns/op	      36 B/op	       1 allocs/op
BenchmarkAdler32-24   	     200	  67560996 ns/op	      16 B/op	       1 allocs/op
BenchmarkCRC32-24     	     100	 105715304 ns/op	      16 B/op	       1 allocs/op
BenchmarkCRC64-24     	      50	 299826700 ns/op	      57 B/op	       1 allocs/op
BenchmarkFNV32-24     	     100	 122715813 ns/op	      16 B/op	       1 allocs/op
BenchmarkFNV32a-24    	     100	 127094626 ns/op	      16 B/op	       1 allocs/op
BenchmarkFNV64-24     	     100	 134332333 ns/op	      16 B/op	       1 allocs/op
BenchmarkFNV64a-24    	     100	 140174186 ns/op	      16 B/op	       1 allocs/op
ok  	github.com/jacobmarble/go-hash-performance	305.137s
```
