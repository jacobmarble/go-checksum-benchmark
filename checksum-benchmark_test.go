package checksumbenchmark

import (
	"crypto/md5"
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

// Test with about 1 MB of data.
const fodderSize int = 1e8

// Test against this.
var fodder []byte = make([]byte, fodderSize)

func checksumTest(b *testing.B, h hash.Hash) {
	for i := 0; i < b.N; i++ {
		h.Write(fodder)
		h.Sum(nil)
	}
}

func BenchmarkMD4(b *testing.B) {
	checksumTest(b, md4.New())
}

func BenchmarkMD5(b *testing.B) {
	checksumTest(b, md5.New())
}

func BenchmarkSHA1(b *testing.B) {
	checksumTest(b, sha1.New())
}

func BenchmarkSHA224(b *testing.B) {
	checksumTest(b, sha256.New224())
}

func BenchmarkSHA256(b *testing.B) {
	checksumTest(b, sha256.New())
}

func BenchmarkSHA384(b *testing.B) {
	checksumTest(b, sha512.New384())
}

func BenchmarkSHA512(b *testing.B) {
	checksumTest(b, sha512.New())
}

func BenchmarkSHA3_224(b *testing.B) {
	checksumTest(b, sha3.New224())
}

func BenchmarkSHA3_256(b *testing.B) {
	checksumTest(b, sha3.New256())
}

func BenchmarkSHA3_384(b *testing.B) {
	checksumTest(b, sha3.New384())
}

func BenchmarkSHA3_512(b *testing.B) {
	checksumTest(b, sha3.New512())
}

func BenchmarkSHA512_224(b *testing.B) {
	checksumTest(b, sha512.New512_224())
}

func BenchmarkSHA512_256(b *testing.B) {
	checksumTest(b, sha512.New512_256())
}

func BenchmarkAdler32(b *testing.B) {
	checksumTest(b, adler32.New())
}

func BenchmarkCRC32(b *testing.B) {
	checksumTest(b, crc32.NewIEEE())
}

func BenchmarkCRC64(b *testing.B) {
	checksumTest(b, crc64.New(crc64.MakeTable(crc64.ISO)))
}

func BenchmarkFNV32(b *testing.B) {
	checksumTest(b, fnv.New32())
}

func BenchmarkFNV32a(b *testing.B) {
	checksumTest(b, fnv.New32a())
}

func BenchmarkFNV64(b *testing.B) {
	checksumTest(b, fnv.New64())
}

func BenchmarkFNV64a(b *testing.B) {
	checksumTest(b, fnv.New64a())
}
