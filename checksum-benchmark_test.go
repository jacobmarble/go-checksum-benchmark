package checksumbenchmark

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"testing"

	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/sha3"
)

func createRandomSlize(size int) []byte {
	s := make([]byte, size, size)
	rand.Read(s)
	return s
}

// Test against this.
var fodder100MB = createRandomSlize(100 * 1024 * 1024)
var fodder1MB = createRandomSlize(1024 * 1024)
var fodder1KB = createRandomSlize(1024)
var fodder64B = createRandomSlize(64)
var fodder24B = createRandomSlize(24)

func checksumTest(b *testing.B, h hash.Hash, fodder []byte) {
	for i := 0; i < b.N; i++ {
		h.Write(fodder)
		h.Sum(nil)
	}
}

// 100 MB benchmarks

func BenchmarkMD4_100MB(b *testing.B) {
	checksumTest(b, md4.New(), fodder100MB)
}

func BenchmarkMD5_100MB(b *testing.B) {
	checksumTest(b, md5.New(), fodder100MB)
}

func BenchmarkSHA1_100MB(b *testing.B) {
	checksumTest(b, sha1.New(), fodder100MB)
}

func BenchmarkSHA224_100MB(b *testing.B) {
	checksumTest(b, sha256.New224(), fodder100MB)
}

func BenchmarkSHA256_100MB(b *testing.B) {
	checksumTest(b, sha256.New(), fodder100MB)
}

func BenchmarkSHA384_100MB(b *testing.B) {
	checksumTest(b, sha512.New384(), fodder100MB)
}

func BenchmarkSHA512_100MB(b *testing.B) {
	checksumTest(b, sha512.New(), fodder100MB)
}

func BenchmarkSHA3_224_100MB(b *testing.B) {
	checksumTest(b, sha3.New224(), fodder100MB)
}

func BenchmarkSHA3_256_100MB(b *testing.B) {
	checksumTest(b, sha3.New256(), fodder100MB)
}

func BenchmarkSHA3_384_100MB(b *testing.B) {
	checksumTest(b, sha3.New384(), fodder100MB)
}

func BenchmarkSHA3_512_100MB(b *testing.B) {
	checksumTest(b, sha3.New512(), fodder100MB)
}

func BenchmarkSHA512_224_100MB(b *testing.B) {
	checksumTest(b, sha512.New512_224(), fodder100MB)
}

func BenchmarkSHA512_256_100MB(b *testing.B) {
	checksumTest(b, sha512.New512_256(), fodder100MB)
}

func BenchmarkAdler32_100MB(b *testing.B) {
	checksumTest(b, adler32.New(), fodder100MB)
}

func BenchmarkCRC32IEEE_100MB(b *testing.B) {
	checksumTest(b, crc32.NewIEEE(), fodder100MB)
}

func BenchmarkCRC32Castagnoli_100MB(b *testing.B) {
	checksumTest(b, crc32.New(crc32.MakeTable(crc32.Castagnoli)), fodder100MB)
}

func BenchmarkCRC32Koopman_100MB(b *testing.B) {
	checksumTest(b, crc32.New(crc32.MakeTable(crc32.Koopman)), fodder100MB)
}

func BenchmarkCRC64_100MB(b *testing.B) {
	checksumTest(b, crc64.New(crc64.MakeTable(crc64.ISO)), fodder100MB)
}

func BenchmarkFNV32_100MB(b *testing.B) {
	checksumTest(b, fnv.New32(), fodder100MB)
}

func BenchmarkFNV32a_100MB(b *testing.B) {
	checksumTest(b, fnv.New32a(), fodder100MB)
}

func BenchmarkFNV64_100MB(b *testing.B) {
	checksumTest(b, fnv.New64(), fodder100MB)
}

func BenchmarkFNV64a_100MB(b *testing.B) {
	checksumTest(b, fnv.New64a(), fodder100MB)
}

// 1 MB benchmarks

func BenchmarkMD4_1MB(b *testing.B) {
	checksumTest(b, md4.New(), fodder1MB)
}

func BenchmarkMD5_1MB(b *testing.B) {
	checksumTest(b, md5.New(), fodder1MB)
}

func BenchmarkSHA1_1MB(b *testing.B) {
	checksumTest(b, sha1.New(), fodder1MB)
}

func BenchmarkSHA224_1MB(b *testing.B) {
	checksumTest(b, sha256.New224(), fodder1MB)
}

func BenchmarkSHA256_1MB(b *testing.B) {
	checksumTest(b, sha256.New(), fodder1MB)
}

func BenchmarkSHA384_1MB(b *testing.B) {
	checksumTest(b, sha512.New384(), fodder1MB)
}

func BenchmarkSHA512_1MB(b *testing.B) {
	checksumTest(b, sha512.New(), fodder1MB)
}

func BenchmarkSHA3_224_1MB(b *testing.B) {
	checksumTest(b, sha3.New224(), fodder1MB)
}

func BenchmarkSHA3_256_1MB(b *testing.B) {
	checksumTest(b, sha3.New256(), fodder1MB)
}

func BenchmarkSHA3_384_1MB(b *testing.B) {
	checksumTest(b, sha3.New384(), fodder1MB)
}

func BenchmarkSHA3_512_1MB(b *testing.B) {
	checksumTest(b, sha3.New512(), fodder1MB)
}

func BenchmarkSHA512_224_1MB(b *testing.B) {
	checksumTest(b, sha512.New512_224(), fodder1MB)
}

func BenchmarkSHA512_256_1MB(b *testing.B) {
	checksumTest(b, sha512.New512_256(), fodder1MB)
}

func BenchmarkAdler32_1MB(b *testing.B) {
	checksumTest(b, adler32.New(), fodder1MB)
}

func BenchmarkCRC32IEEE_1MB(b *testing.B) {
	checksumTest(b, crc32.NewIEEE(), fodder1MB)
}

func BenchmarkCRC32Castagnoli_1MB(b *testing.B) {
	checksumTest(b, crc32.New(crc32.MakeTable(crc32.Castagnoli)), fodder1MB)
}

func BenchmarkCRC32Koopman_1MB(b *testing.B) {
	checksumTest(b, crc32.New(crc32.MakeTable(crc32.Koopman)), fodder1MB)
}

func BenchmarkCRC64_1MB(b *testing.B) {
	checksumTest(b, crc64.New(crc64.MakeTable(crc64.ISO)), fodder1MB)
}

func BenchmarkFNV32_1MB(b *testing.B) {
	checksumTest(b, fnv.New32(), fodder1MB)
}

func BenchmarkFNV32a_1MB(b *testing.B) {
	checksumTest(b, fnv.New32a(), fodder1MB)
}

func BenchmarkFNV64_1MB(b *testing.B) {
	checksumTest(b, fnv.New64(), fodder1MB)
}

func BenchmarkFNV64a_1MB(b *testing.B) {
	checksumTest(b, fnv.New64a(), fodder1MB)
}

// 1KB benchmarks

func BenchmarkMD4_1KB(b *testing.B) {
	checksumTest(b, md4.New(), fodder1KB)
}

func BenchmarkMD5_1KB(b *testing.B) {
	checksumTest(b, md5.New(), fodder1KB)
}

func BenchmarkSHA1_1KB(b *testing.B) {
	checksumTest(b, sha1.New(), fodder1KB)
}

func BenchmarkSHA224_1KB(b *testing.B) {
	checksumTest(b, sha256.New224(), fodder1KB)
}

func BenchmarkSHA256_1KB(b *testing.B) {
	checksumTest(b, sha256.New(), fodder1KB)
}

func BenchmarkSHA384_1KB(b *testing.B) {
	checksumTest(b, sha512.New384(), fodder1KB)
}

func BenchmarkSHA512_1KB(b *testing.B) {
	checksumTest(b, sha512.New(), fodder1KB)
}

func BenchmarkSHA3_224_1KB(b *testing.B) {
	checksumTest(b, sha3.New224(), fodder1KB)
}

func BenchmarkSHA3_256_1KB(b *testing.B) {
	checksumTest(b, sha3.New256(), fodder1KB)
}

func BenchmarkSHA3_384_1KB(b *testing.B) {
	checksumTest(b, sha3.New384(), fodder1KB)
}

func BenchmarkSHA3_512_1KB(b *testing.B) {
	checksumTest(b, sha3.New512(), fodder1KB)
}

func BenchmarkSHA512_224_1KB(b *testing.B) {
	checksumTest(b, sha512.New512_224(), fodder1KB)
}

func BenchmarkSHA512_256_1KB(b *testing.B) {
	checksumTest(b, sha512.New512_256(), fodder1KB)
}

func BenchmarkAdler32_1KB(b *testing.B) {
	checksumTest(b, adler32.New(), fodder1KB)
}

func BenchmarkCRC32IEEE_1KB(b *testing.B) {
	checksumTest(b, crc32.NewIEEE(), fodder1KB)
}

func BenchmarkCRC32Castagnoli_1KB(b *testing.B) {
	checksumTest(b, crc32.New(crc32.MakeTable(crc32.Castagnoli)), fodder1KB)
}

func BenchmarkCRC32Koopman_1KB(b *testing.B) {
	checksumTest(b, crc32.New(crc32.MakeTable(crc32.Koopman)), fodder1KB)
}

func BenchmarkCRC64_1KB(b *testing.B) {
	checksumTest(b, crc64.New(crc64.MakeTable(crc64.ISO)), fodder1KB)
}

func BenchmarkFNV32_1KB(b *testing.B) {
	checksumTest(b, fnv.New32(), fodder1KB)
}

func BenchmarkFNV32a_1KB(b *testing.B) {
	checksumTest(b, fnv.New32a(), fodder1KB)
}

func BenchmarkFNV64_1KB(b *testing.B) {
	checksumTest(b, fnv.New64(), fodder1KB)
}

func BenchmarkFNV64a_1KB(b *testing.B) {
	checksumTest(b, fnv.New64a(), fodder1KB)
}

// 64B benchmarks

func BenchmarkMD4_64B(b *testing.B) {
	checksumTest(b, md4.New(), fodder64B)
}

func BenchmarkMD5_64B(b *testing.B) {
	checksumTest(b, md5.New(), fodder64B)
}

func BenchmarkSHA1_64B(b *testing.B) {
	checksumTest(b, sha1.New(), fodder64B)
}

func BenchmarkSHA224_64B(b *testing.B) {
	checksumTest(b, sha256.New224(), fodder64B)
}

func BenchmarkSHA256_64B(b *testing.B) {
	checksumTest(b, sha256.New(), fodder64B)
}

func BenchmarkSHA384_64B(b *testing.B) {
	checksumTest(b, sha512.New384(), fodder64B)
}

func BenchmarkSHA512_64B(b *testing.B) {
	checksumTest(b, sha512.New(), fodder64B)
}

func BenchmarkSHA3_224_64B(b *testing.B) {
	checksumTest(b, sha3.New224(), fodder64B)
}

func BenchmarkSHA3_256_64B(b *testing.B) {
	checksumTest(b, sha3.New256(), fodder64B)
}

func BenchmarkSHA3_384_64B(b *testing.B) {
	checksumTest(b, sha3.New384(), fodder64B)
}

func BenchmarkSHA3_512_64B(b *testing.B) {
	checksumTest(b, sha3.New512(), fodder64B)
}

func BenchmarkSHA512_224_64B(b *testing.B) {
	checksumTest(b, sha512.New512_224(), fodder64B)
}

func BenchmarkSHA512_256_64B(b *testing.B) {
	checksumTest(b, sha512.New512_256(), fodder64B)
}

func BenchmarkAdler32_64B(b *testing.B) {
	checksumTest(b, adler32.New(), fodder64B)
}

func BenchmarkCRC32IEEE_64B(b *testing.B) {
	checksumTest(b, crc32.NewIEEE(), fodder64B)
}

func BenchmarkCRC32Castagnoli_64B(b *testing.B) {
	checksumTest(b, crc32.New(crc32.MakeTable(crc32.Castagnoli)), fodder64B)
}

func BenchmarkCRC32Koopman_64B(b *testing.B) {
	checksumTest(b, crc32.New(crc32.MakeTable(crc32.Koopman)), fodder64B)
}

func BenchmarkCRC64_64B(b *testing.B) {
	checksumTest(b, crc64.New(crc64.MakeTable(crc64.ISO)), fodder64B)
}

func BenchmarkFNV32_64B(b *testing.B) {
	checksumTest(b, fnv.New32(), fodder64B)
}

func BenchmarkFNV32a_64B(b *testing.B) {
	checksumTest(b, fnv.New32a(), fodder64B)
}

func BenchmarkFNV64_64B(b *testing.B) {
	checksumTest(b, fnv.New64(), fodder64B)
}

func BenchmarkFNV64a_64B(b *testing.B) {
	checksumTest(b, fnv.New64a(), fodder64B)
}

// 24B benchmarks

func BenchmarkMD4_24B(b *testing.B) {
	checksumTest(b, md4.New(), fodder24B)
}

func BenchmarkMD5_24B(b *testing.B) {
	checksumTest(b, md5.New(), fodder24B)
}

func BenchmarkSHA1_24B(b *testing.B) {
	checksumTest(b, sha1.New(), fodder24B)
}

func BenchmarkSHA224_24B(b *testing.B) {
	checksumTest(b, sha256.New224(), fodder24B)
}

func BenchmarkSHA256_24B(b *testing.B) {
	checksumTest(b, sha256.New(), fodder24B)
}

func BenchmarkSHA384_24B(b *testing.B) {
	checksumTest(b, sha512.New384(), fodder24B)
}

func BenchmarkSHA512_24B(b *testing.B) {
	checksumTest(b, sha512.New(), fodder24B)
}

func BenchmarkSHA3_224_24B(b *testing.B) {
	checksumTest(b, sha3.New224(), fodder24B)
}

func BenchmarkSHA3_256_24B(b *testing.B) {
	checksumTest(b, sha3.New256(), fodder24B)
}

func BenchmarkSHA3_384_24B(b *testing.B) {
	checksumTest(b, sha3.New384(), fodder24B)
}

func BenchmarkSHA3_512_24B(b *testing.B) {
	checksumTest(b, sha3.New512(), fodder24B)
}

func BenchmarkSHA512_224_24B(b *testing.B) {
	checksumTest(b, sha512.New512_224(), fodder24B)
}

func BenchmarkSHA512_256_24B(b *testing.B) {
	checksumTest(b, sha512.New512_256(), fodder24B)
}

func BenchmarkAdler32_24B(b *testing.B) {
	checksumTest(b, adler32.New(), fodder24B)
}

func BenchmarkCRC32IEEE_24B(b *testing.B) {
	checksumTest(b, crc32.NewIEEE(), fodder24B)
}

func BenchmarkCRC32Castagnoli_24B(b *testing.B) {
	checksumTest(b, crc32.New(crc32.MakeTable(crc32.Castagnoli)), fodder24B)
}

func BenchmarkCRC32Koopman_24B(b *testing.B) {
	checksumTest(b, crc32.New(crc32.MakeTable(crc32.Koopman)), fodder24B)
}

func BenchmarkCRC64_24B(b *testing.B) {
	checksumTest(b, crc64.New(crc64.MakeTable(crc64.ISO)), fodder24B)
}

func BenchmarkFNV32_24B(b *testing.B) {
	checksumTest(b, fnv.New32(), fodder24B)
}

func BenchmarkFNV32a_24B(b *testing.B) {
	checksumTest(b, fnv.New32a(), fodder24B)
}

func BenchmarkFNV64_24B(b *testing.B) {
	checksumTest(b, fnv.New64(), fodder24B)
}

func BenchmarkFNV64a_24B(b *testing.B) {
	checksumTest(b, fnv.New64a(), fodder24B)
}
