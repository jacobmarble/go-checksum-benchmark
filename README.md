I needed a fast (not necessarily cryptogrphic) checksum algorithm.
Some complete results are in [the results directory](results).
Typical results:

```console
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/jacobmarble/go-checksum-benchmark
cpu: Intel(R) Xeon(R) CPU E5-2670 0 @ 2.60GHz
BenchmarkChecksums_Cryptographic/md4-1B-16                      1000000       1189   ns/op     120 B/op   3 allocs/op
BenchmarkChecksums_Cryptographic/md5-1B-16                      1974649        602.7 ns/op     112 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/sha1-1B-16                     1720981        854.9 ns/op     136 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/sha256-1B-16                    741099       1498   ns/op     160 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/sha512-1B-16                    573054       1925   ns/op     288 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/sha3-256-1B-16                  325473       3273   ns/op     960 B/op   4 allocs/op
BenchmarkChecksums_Cryptographic/sha3-512-1B-16                  332496       3206   ns/op    1024 B/op   4 allocs/op
BenchmarkChecksums_Cryptographic/sha512-224-1B-16                577180       1920   ns/op     256 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/sha512-256-1B-16                565468       1913   ns/op     256 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/siphash-1B-16                  3360697        357.2 ns/op      88 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/blake2b-256-1B-16              1205532        994.1 ns/op     416 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/blake2b-512-1B-16              1184875       1013   ns/op     448 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/blake3-1B-16                    280167       4157   ns/op   10944 B/op   3 allocs/op
BenchmarkChecksums_Cryptographic/md4-1KB-16                      231244       5297   ns/op     120 B/op   3 allocs/op
BenchmarkChecksums_Cryptographic/md5-1KB-16                      338662       3337   ns/op     112 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/sha1-1KB-16                     219108       5052   ns/op     136 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/sha256-1KB-16                   118611      10255   ns/op     160 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/sha512-1KB-16                   122155       9051   ns/op     288 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/sha3-256-1KB-16                  90908      13206   ns/op     960 B/op   4 allocs/op
BenchmarkChecksums_Cryptographic/sha3-512-1KB-16                  60944      21019   ns/op    1024 B/op   4 allocs/op
BenchmarkChecksums_Cryptographic/sha512-224-1KB-16               153292       8131   ns/op     256 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/sha512-256-1KB-16               144727       8265   ns/op     256 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/siphash-1KB-16                  732279       1909   ns/op      88 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/blake2b-256-1KB-16              248917       4489   ns/op     416 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/blake2b-512-1KB-16              247387       4514   ns/op     448 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/blake3-1KB-16                   159715       7405   ns/op   10944 B/op   3 allocs/op
BenchmarkChecksums_Cryptographic/md4-1MB-16                         284    3984995   ns/op     120 B/op   3 allocs/op
BenchmarkChecksums_Cryptographic/md5-1MB-16                         470    2249365   ns/op     112 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/sha1-1MB-16                        320    3308465   ns/op     136 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/sha256-1MB-16                      134    8563483   ns/op     160 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/sha512-1MB-16                      222    5459345   ns/op     288 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/sha3-256-1MB-16                    207    5510410   ns/op     960 B/op   4 allocs/op
BenchmarkChecksums_Cryptographic/sha3-512-1MB-16                    100   10410902   ns/op    1024 B/op   4 allocs/op
BenchmarkChecksums_Cryptographic/sha512-224-1MB-16                  224    5525164   ns/op     256 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/sha512-256-1MB-16                  207    5511269   ns/op     256 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/siphash-1MB-16                    1315     911737   ns/op      88 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/blake2b-256-1MB-16                 667    1834133   ns/op     416 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/blake2b-512-1MB-16                 566    1778802   ns/op     448 B/op   2 allocs/op
BenchmarkChecksums_Cryptographic/blake3-1MB-16                      550    2012875   ns/op   10944 B/op   3 allocs/op
BenchmarkChecksums_NonCryptographic/adler32-1B-16               5871230        189.3 ns/op      16 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/crc32-ieee-1B-16            4116416        275.4 ns/op      24 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/crc32-castagnoli-1B-16      4641223        284.5 ns/op      24 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/crc32-koopman-1B-16          101482      11325   ns/op    1048 B/op   3 allocs/op
BenchmarkChecksums_NonCryptographic/crc64-1B-16                 4423560        271.8 ns/op      24 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv32-1B-16                 6000358        195.6 ns/op      16 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv32a-1B-16                6390013        182.8 ns/op      16 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv64-1B-16                 6302191        186.5 ns/op      16 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv64a-1B-16                6451806        182.8 ns/op      16 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv128-1B-16                5411886        219.1 ns/op      32 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv128a-1B-16               5333749        222.2 ns/op      32 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxh3-1B-16                  5334547        221.2 ns/op      32 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxh3-128-1B-16              4176370        286.6 ns/op      48 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxhash-cespare-1B-16        4689472        250.5 ns/op      88 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxhash-oneofone-32-1B-16    4929764        240.8 ns/op      56 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxhash-oneofone-64-1B-16    4436134        267.2 ns/op     104 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxhash-pierrec-32-1B-16     4273500        279.1 ns/op      72 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxhash-pierrec-64-1B-16     3754924        314.0 ns/op     104 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/adler32-1KB-16              1252210        892.5 ns/op      16 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/crc32-ieee-1KB-16           1422889        829.5 ns/op      24 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/crc32-castagnoli-1KB-16     2748855        402.8 ns/op      24 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/crc32-koopman-1KB-16          67281      17630   ns/op    1048 B/op   3 allocs/op
BenchmarkChecksums_NonCryptographic/crc64-1KB-16                 880332       1426   ns/op      24 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv32-1KB-16                 641563       1799   ns/op      16 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv32a-1KB-16                641605       1849   ns/op      16 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv64-1KB-16                 630285       1728   ns/op      16 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv64a-1KB-16                642220       1827   ns/op      16 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv128-1KB-16                255008       4457   ns/op      32 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv128a-1KB-16               241185       4684   ns/op      32 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxh3-1KB-16                 2998538        413.0 ns/op      32 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxh3-128-1KB-16             2395042        496.7 ns/op      48 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxhash-cespare-1KB-16       2307170        515.0 ns/op      88 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxhash-oneofone-32-1KB-16    951399       1062   ns/op      56 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxhash-oneofone-64-1KB-16   1952906        615.4 ns/op     104 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxhash-pierrec-32-1KB-16    1207176        965.1 ns/op      72 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxhash-pierrec-64-1KB-16    1667088        714.1 ns/op     104 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/adler32-1MB-16                 1731     694490   ns/op      16 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/crc32-ieee-1MB-16              2529     458401   ns/op      24 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/crc32-castagnoli-1MB-16       18445      62310   ns/op      24 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/crc32-koopman-1MB-16            322    3281120   ns/op    1048 B/op   3 allocs/op
BenchmarkChecksums_NonCryptographic/crc64-1MB-16                   1088    1091495   ns/op      24 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv32-1MB-16                    739    1627413   ns/op      16 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv32a-1MB-16                   638    1623955   ns/op      16 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv64-1MB-16                    655    1678922   ns/op      16 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv64a-1MB-16                   729    1619218   ns/op      16 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv128-1MB-16                   262    4275677   ns/op      32 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/fnv128a-1MB-16                  260    4320797   ns/op      32 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxh3-1MB-16                   13340      88124   ns/op      32 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxh3-128-1MB-16               13036      88590   ns/op      48 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxhash-cespare-1MB-16         10348     116291   ns/op      88 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxhash-oneofone-32-1MB-16      2581     469765   ns/op      56 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxhash-oneofone-64-1MB-16      7368     153985   ns/op     104 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxhash-pierrec-32-1MB-16       3607     344126   ns/op      72 B/op   2 allocs/op
BenchmarkChecksums_NonCryptographic/xxhash-pierrec-64-1MB-16       6283     181858   ns/op     104 B/op   2 allocs/op
PASS
ok      github.com/jacobmarble/go-checksum-benchmark    150.689s
```
