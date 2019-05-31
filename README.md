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

# Benchmarks for different CPU and OS

### CPU Core i7-7700 @ 3.60 GHz
go version: go1.12.3 windows/amd64

OS: Windows 10

```
go test -bench . -benchtime 10s -benchmem
Hash = b68ced3f33a4533106a5fc1195e4991c3b85f052d4c2fd2dd3a5ac3443f319c3
goos: windows
goarch: amd64
pkg: gitlab.kazdream.kz/kronos/cdr-lldb
BenchmarkMD4-8                        50         290930338 ns/op              25 B/op          2 allocs/op
BenchmarkMD5-8                       100         135470023 ns/op              16 B/op          1 allocs/op
BenchmarkSHA1-8                      200          97099841 ns/op              32 B/op          1 allocs/op
BenchmarkSHA224-8                     50         238879912 ns/op              34 B/op          1 allocs/op
BenchmarkSHA256-8                     50         237999942 ns/op              34 B/op          1 allocs/op
BenchmarkSHA384-8                    100         163843931 ns/op              50 B/op          1 allocs/op
BenchmarkSHA512-8                    100         159730001 ns/op              66 B/op          1 allocs/op
BenchmarkSHA3_224-8                   50         271539850 ns/op             520 B/op          3 allocs/op
BenchmarkSHA3_256-8                   50         286780158 ns/op             520 B/op          3 allocs/op
BenchmarkSHA3_384-8                   50         373699702 ns/op             552 B/op          3 allocs/op
BenchmarkSHA3_512-8                   30         540801056 ns/op             590 B/op          3 allocs/op
BenchmarkSHA512_224-8                100         160399985 ns/op              34 B/op          1 allocs/op
BenchmarkSHA512_256-8                100         160199961 ns/op              34 B/op          1 allocs/op
BenchmarkAdler32-8                   500          36528002 ns/op               8 B/op          1 allocs/op
BenchmarkCRC32-8                    2000           6639999 ns/op               8 B/op          1 allocs/op
BenchmarkCRC64-8                     300          53929997 ns/op               8 B/op          1 allocs/op
BenchmarkFNV32-8                     100         107619997 ns/op               8 B/op          1 allocs/op
BenchmarkFNV32a-8                    100         107689756 ns/op               8 B/op          1 allocs/op
BenchmarkFNV64-8                     100         107600193 ns/op               8 B/op          1 allocs/op
BenchmarkFNV64a-8                    100         107639708 ns/op               8 B/op          1 allocs/op
PASS
```

### Intel(R) Xeon(R) CPU E5-2620 v4 @ 2.10GHz
OS: CentOS 7

go version: go1.11 linux/amd64

```
go test -bench . -benchtime 10s -benchmem
goos: linux
goarch: amd64
pkg: hashbench
BenchmarkMD4-32                       20         706462031 ns/op             112 B/op          2 allocs/op
BenchmarkMD5-32                       50         231929740 ns/op              51 B/op          1 allocs/op
BenchmarkSHA1-32                     100         156139352 ns/op              49 B/op          1 allocs/op
BenchmarkSHA224-32                    30         546395248 ns/op              36 B/op          1 allocs/op
BenchmarkSHA256-32                    30         399602223 ns/op              36 B/op          1 allocs/op
BenchmarkSHA384-32                    50         291316968 ns/op              52 B/op          1 allocs/op
BenchmarkSHA512-32                    50         275933078 ns/op              68 B/op          1 allocs/op
BenchmarkSHA3_224-32                  30         464524838 ns/op             526 B/op          3 allocs/op
BenchmarkSHA3_256-32                  30         510971514 ns/op             526 B/op          3 allocs/op
BenchmarkSHA3_384-32                  20         624208137 ns/op             566 B/op          3 allocs/op
BenchmarkSHA3_512-32                  20        1139384465 ns/op             598 B/op          3 allocs/op
BenchmarkSHA512_224-32                50         278495297 ns/op              36 B/op          1 allocs/op
BenchmarkSHA512_256-32                50         293022377 ns/op              36 B/op          1 allocs/op
BenchmarkAdler32-32                  200          71743513 ns/op               8 B/op          1 allocs/op
BenchmarkCRC32-32                   2000           7926195 ns/op               8 B/op          1 allocs/op
BenchmarkCRC64-32                    200         108477913 ns/op               8 B/op          1 allocs/op
BenchmarkFNV32-32                    100         187123349 ns/op               8 B/op          1 allocs/op
BenchmarkFNV32a-32                   100         176686089 ns/op               8 B/op          1 allocs/op
BenchmarkFNV64-32                    100         178755620 ns/op               8 B/op          1 allocs/op
BenchmarkFNV64a-32                   100         174152024 ns/op               8 B/op          1 allocs/op
PASS
ok      hashbench       403.514s
```

### Intel(R) Xeon(R) Gold 6130 CPU @ 2.10GHz
OS: Ubuntu 16.04 LTS

go version: go1.12.5 linux/amd64
```
go test -bench . -benchtime 10s -benchmem
goos: linux
goarch: amd64
pkg: hashbench
BenchmarkMD4-64                       50         307084249 ns/op              25 B/op          2 allocs/op
BenchmarkMD5-64                      100         143734794 ns/op              16 B/op          1 allocs/op
BenchmarkSHA1-64                     200         100104403 ns/op              32 B/op          1 allocs/op
BenchmarkSHA224-64                    50         255164959 ns/op              34 B/op          1 allocs/op
BenchmarkSHA256-64                    50         252682722 ns/op              34 B/op          1 allocs/op
BenchmarkSHA384-64                   100         171539314 ns/op              50 B/op          1 allocs/op
BenchmarkSHA512-64                   100         172867683 ns/op              66 B/op          1 allocs/op
BenchmarkSHA3_224-64                  50         291099656 ns/op             520 B/op          3 allocs/op
BenchmarkSHA3_256-64                  50         308241885 ns/op             520 B/op          3 allocs/op
BenchmarkSHA3_384-64                  30         399775618 ns/op             558 B/op          3 allocs/op
BenchmarkSHA3_512-64                  20         577732120 ns/op             598 B/op          3 allocs/op
BenchmarkSHA512_224-64               100         173954527 ns/op              34 B/op          1 allocs/op
BenchmarkSHA512_256-64               100         172649419 ns/op              34 B/op          1 allocs/op
BenchmarkAdler32-64                  300          39711318 ns/op               8 B/op          1 allocs/op
BenchmarkCRC32-64                   3000           5085715 ns/op               8 B/op          1 allocs/op
BenchmarkCRC64-64                    300          58425119 ns/op               8 B/op          1 allocs/op
BenchmarkFNV32-64                    100         118456634 ns/op               8 B/op          1 allocs/op
BenchmarkFNV32a-64                   100         119094280 ns/op               8 B/op          1 allocs/op
BenchmarkFNV64-64                    100         120866834 ns/op               8 B/op          1 allocs/op
BenchmarkFNV64a-64                   100         117941566 ns/op               8 B/op          1 allocs/op
PASS
ok      hashbench       314.485s
```
