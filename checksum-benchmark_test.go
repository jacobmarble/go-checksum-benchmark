package checksumbenchmark

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"math/rand"
	"regexp"
	"strconv"
	"testing"

	xxhashoneofone "github.com/OneOfOne/xxhash"
	xxhashcespare "github.com/cespare/xxhash/v2"
	xxhashpierrec32 "github.com/pierrec/xxHash/xxHash32"
	xxhashpierrec64 "github.com/pierrec/xxHash/xxHash64"
	"github.com/zeebo/xxh3"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/sha3"
)

func BenchmarkChecksums(b *testing.B) {
	for _, size := range []string{"8B", "256B", "8KB", "256KB", "8MB"} {
		fodder := generateRandomBytes(b, size)
		for _, hi := range hashImplementations {
			h := hi.hashNewFunc()
			b.Run(fmt.Sprintf("%s-%sytes", hi.hashName, size), func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					h.Write(fodder)
					h.Sum(nil)
					h.Reset()
				}
			})
		}
	}
}

var hashImplementations = []struct {
	hashName    string
	hashNewFunc func() hash.Hash
}{{
	hashName:    "md4",
	hashNewFunc: md4.New,
}, {
	hashName:    "md5",
	hashNewFunc: md5.New,
}, {
	hashName:    "sha1",
	hashNewFunc: sha1.New,
}, {
	hashName:    "sha224",
	hashNewFunc: sha256.New224,
}, {
	hashName:    "sha256",
	hashNewFunc: sha256.New,
}, {
	hashName:    "sha384",
	hashNewFunc: sha512.New384,
}, {
	hashName:    "sha512",
	hashNewFunc: sha512.New,
}, {
	hashName:    "sha3-224",
	hashNewFunc: sha3.New224,
}, {
	hashName:    "sha3-256",
	hashNewFunc: sha3.New256,
}, {
	hashName:    "sha3-384",
	hashNewFunc: sha3.New384,
}, {
	hashName:    "sha3-512",
	hashNewFunc: sha3.New512,
}, {
	hashName:    "sha512-224",
	hashNewFunc: sha512.New512_224,
}, {
	hashName:    "sha512-256",
	hashNewFunc: sha512.New512_256,
}, {
	hashName:    "adler32",
	hashNewFunc: func() hash.Hash { return adler32.New() },
}, {
	hashName:    "crc32-ieee",
	hashNewFunc: func() hash.Hash { return crc32.NewIEEE() },
}, {
	hashName:    "crc32-castagnoli",
	hashNewFunc: func() hash.Hash { return crc32.New(crc32.MakeTable(crc32.Castagnoli)) },
}, {
	hashName:    "crc32-koopman",
	hashNewFunc: func() hash.Hash { return crc32.New(crc32.MakeTable(crc32.Koopman)) },
}, {
	hashName:    "crc64",
	hashNewFunc: func() hash.Hash { return crc64.New(crc64.MakeTable(crc64.ISO)) },
}, {
	hashName:    "fnv32",
	hashNewFunc: func() hash.Hash { return fnv.New32() },
}, {
	hashName:    "fnv32a",
	hashNewFunc: func() hash.Hash { return fnv.New32a() },
}, {
	hashName:    "fnv64",
	hashNewFunc: func() hash.Hash { return fnv.New64() },
}, {
	hashName:    "fnv64a",
	hashNewFunc: func() hash.Hash { return fnv.New64a() },
}, {
	hashName:    "xxh3",
	hashNewFunc: func() hash.Hash { return newXXH3Wrapper() },
}, {
	hashName:    "xxh3-128",
	hashNewFunc: func() hash.Hash { return newXXH3128Wrapper() },
}, {
	hashName:    "xxhash-cespare",
	hashNewFunc: func() hash.Hash { return xxhashcespare.New() },
}, {
	hashName:    "xxhash-oneofone-32",
	hashNewFunc: func() hash.Hash { return xxhashoneofone.New32() },
}, {
	hashName:    "xxhash-oneofone-64",
	hashNewFunc: func() hash.Hash { return xxhashoneofone.New64() },
}, {
	hashName:    "xxhash-pierrec-32",
	hashNewFunc: func() hash.Hash { return xxhashpierrec32.New(0xCAFE) },
}, {
	hashName:    "xxhash-pierrec-64",
	hashNewFunc: func() hash.Hash { return xxhashpierrec64.New(0xCAFE) },
}}

var stringToSize = regexp.MustCompile(`(\d)([KMG]?)B`)

func generateRandomBytes(tb testing.TB, sizeString string) []byte {
	ss := stringToSize.FindStringSubmatch(sizeString)
	if len(ss) != 3 {
		tb.Fatalf("invalid size string %q", ss)
	}
	size, err := strconv.Atoi(ss[1])
	if err != nil {
		tb.Fatalf("failed to parse integer %q", ss[1])
	}
	switch ss[2] {
	case "G":
		size *= 1024
		fallthrough
	case "M":
		size *= 1024
		fallthrough
	case "K":
		size *= 1024
	}
	s := make([]byte, size, size)
	_, _ = rand.Read(s)
	return s
}

type xxh3Wrapper struct {
	hash.Hash
	sum uint64
}

func newXXH3Wrapper() hash.Hash {
	return &xxh3Wrapper{
		sum: xxh3.Hash(nil),
	}
}

func (x *xxh3Wrapper) Write(p []byte) (n int, err error) {
	x.sum = xxh3.Hash(p)
	return len(p), nil
}

func (x *xxh3Wrapper) Sum(b []byte) []byte {
	if len(b) > 0 {
		panic("this implementation only handles Sum(nil)")
	}
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], x.sum)
	return buf[:]
}

func (x *xxh3Wrapper) Reset() {
}

type xxh3128Wrapper struct {
	hash.Hash
	sum [2]uint64
}

func newXXH3128Wrapper() hash.Hash {
	return &xxh3128Wrapper{
		sum: xxh3.Hash128(nil),
	}
}

func (x *xxh3128Wrapper) Write(p []byte) (n int, err error) {
	x.sum = xxh3.Hash128(p)
	return len(p), nil
}

func (x *xxh3128Wrapper) Sum(b []byte) []byte {
	if len(b) > 0 {
		panic("this implementation only handles Sum(nil)")
	}
	var buf [16]byte
	binary.BigEndian.PutUint64(buf[:8], x.sum[0])
	binary.BigEndian.PutUint64(buf[8:], x.sum[1])
	return buf[:]
}

func (x *xxh3128Wrapper) Reset() {
}
