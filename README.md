I needed a fast, not necessarily crypto-secure, checksum algorithm.
Some complete results are in [the results directory](results).
A snippet of typical results follows:

```text
goos: darwin
goarch: amd64
pkg: github.com/jacobmarble/go-checksum-benchmark
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkChecksums/md4-8Bytes-16                  40001572        299.0 ns/op      24 B/op       2 allocs/op
BenchmarkChecksums/md5-8Bytes-16                  97674513        123.5 ns/op      16 B/op       1 allocs/op
BenchmarkChecksums/sha1-8Bytes-16                 74332724        152.6 ns/op      24 B/op       1 allocs/op
BenchmarkChecksums/sha224-8Bytes-16               56857090        216.2 ns/op      32 B/op       1 allocs/op
BenchmarkChecksums/sha256-8Bytes-16               55034511        214.8 ns/op      32 B/op       1 allocs/op
BenchmarkChecksums/sha384-8Bytes-16               44043376        276.2 ns/op      48 B/op       1 allocs/op
BenchmarkChecksums/sha512-8Bytes-16               43375620        281.3 ns/op      64 B/op       1 allocs/op
BenchmarkChecksums/sha3-224-8Bytes-16             19225023        626.5 ns/op     512 B/op       3 allocs/op
BenchmarkChecksums/sha3-256-8Bytes-16             19200603        623.3 ns/op     512 B/op       3 allocs/op
BenchmarkChecksums/sha3-384-8Bytes-16             19492800        608.7 ns/op     544 B/op       3 allocs/op
BenchmarkChecksums/sha3-512-8Bytes-16             20397945        585.1 ns/op     576 B/op       3 allocs/op
BenchmarkChecksums/sha512-224-8Bytes-16           44136050        274.5 ns/op      32 B/op       1 allocs/op
BenchmarkChecksums/sha512-256-8Bytes-16           43263284        272.4 ns/op      32 B/op       1 allocs/op
BenchmarkChecksums/adler32-8Bytes-16             409159914        29.38 ns/op       8 B/op       1 allocs/op
BenchmarkChecksums/crc32-ieee-8Bytes-16          309765574        38.93 ns/op       8 B/op       1 allocs/op
BenchmarkChecksums/crc32-castagnoli-8Bytes-16    444330688        27.38 ns/op       8 B/op       1 allocs/op
BenchmarkChecksums/crc32-koopman-8Bytes-16       354338702        33.80 ns/op       8 B/op       1 allocs/op
BenchmarkChecksums/crc64-8Bytes-16               346713039        34.74 ns/op       8 B/op       1 allocs/op
BenchmarkChecksums/fnv32-8Bytes-16               493735369        24.24 ns/op       8 B/op       1 allocs/op
BenchmarkChecksums/fnv32a-8Bytes-16              495636584        24.62 ns/op       8 B/op       1 allocs/op
BenchmarkChecksums/fnv64-8Bytes-16               499863522        24.08 ns/op       8 B/op       1 allocs/op
BenchmarkChecksums/fnv64a-8Bytes-16              491249101        24.28 ns/op       8 B/op       1 allocs/op
BenchmarkChecksums/xxh3-8Bytes-16                611885392        19.68 ns/op       8 B/op       1 allocs/op
BenchmarkChecksums/xxh3-128-8Bytes-16            404731935        30.33 ns/op      16 B/op       1 allocs/op
BenchmarkChecksums/xxhash-cespare-8Bytes-16      387838120        30.95 ns/op       8 B/op       1 allocs/op
BenchmarkChecksums/xxhash-oneofone-32-8Bytes-16  382575966        31.10 ns/op       8 B/op       1 allocs/op
BenchmarkChecksums/xxhash-oneofone-64-8Bytes-16  386133150        31.07 ns/op       8 B/op       1 allocs/op
BenchmarkChecksums/xxhash-pierrec-32-8Bytes-16   331654134        36.12 ns/op       8 B/op       1 allocs/op
BenchmarkChecksums/xxhash-pierrec-64-8Bytes-16   320278574        37.39 ns/op       8 B/op       1 allocs/op
```
