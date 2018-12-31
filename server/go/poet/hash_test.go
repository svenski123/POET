package poet

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/svenski123/sha256-simd"
	spacemeshos "github.com/spacemeshos/sha256-simd"
	miniosha "github.com/minio/sha256-simd"
)

const n_sha256_trials int64 = 2000000

func BenchmarkSha256(t *testing.B) {
	const prefix = "[minio] "

	buff := bytes.Buffer{}
	buff.Write([]byte("Seed data goes here"))
	out := [32]byte{}

	fmt.Printf("%sComputing %d serial sha-256...\n", prefix, n_sha256_trials)

	t1 := time.Now().UnixNano()

	for i := int64(0); i < n_sha256_trials; i++ {
		out = miniosha.Sum256(buff.Bytes())
		buff.Reset()
		buff.Write(out[:])
	}

	d := float64(time.Now().UnixNano() - t1) / 1e9
	r := float64(n_sha256_trials) / d

	fmt.Printf("%s  Final hash: %x\n", prefix, buff.Bytes())
	fmt.Printf("%sRunning time: %g secs.\n", prefix, d)
	fmt.Printf("%s   Hash-rate: %g hashes-per-sec\n", prefix, r)
}

func BenchmarkSha256Ex(t *testing.B) {
	const prefix = "[space] "

	buff := bytes.Buffer{}
	buff.Write([]byte("Seed data goes here"))
	out := [32]byte{}

	fmt.Printf("%sComputing %d serial sha-256...\n", prefix, n_sha256_trials)

	t1 := time.Now().UnixNano()

	for i := int64(0); i < n_sha256_trials; i++ {
		out = spacemeshos.Sum256(buff.Bytes())
		buff.Reset()
		buff.Write(out[:])
	}

	d := float64(time.Now().UnixNano() - t1) / 1e9
	r := float64(n_sha256_trials) / d

	fmt.Printf("%s  Final hash: %x\n", prefix, buff.Bytes())
	fmt.Printf("%sRunning time: %g secs.\n", prefix, d)
	fmt.Printf("%s   Hash-rate: %g hashes-per-sec\n", prefix, r)
}

func BenchmarkSha256Opt(t *testing.B) {
	const prefix = "[optim] "

	buff := bytes.Buffer{}
	buff.Write([]byte("Seed data goes here"))
	out := [32]byte{}

	fmt.Printf("%sComputing %d serial sha-256...\n", prefix, n_sha256_trials)

	t1 := time.Now().UnixNano()

	for i := int64(0); i < n_sha256_trials; i++ {
		out = sha256.Sum256(buff.Bytes())
		buff.Reset()
		buff.Write(out[:])
	}

	d := float64(time.Now().UnixNano() - t1) / 1e9
	r := float64(n_sha256_trials) / d

	fmt.Printf("%s  Final hash: %x\n", prefix, buff.Bytes())
	fmt.Printf("%sRunning time: %g secs.\n", prefix, d)
	fmt.Printf("%s   Hash-rate: %g hashes-per-sec\n", prefix, r)
}

func BenchmarkSha256HashVals(t *testing.B) {
	const prefix = "[hvals] "

	h := NewSHA256()
	buff := bytes.Buffer{}
	buff.Write([]byte("Seed data goes here"))
	var out []byte;

	fmt.Printf("%sComputing %d serial sha-256...\n", prefix, n_sha256_trials)

	t1 := time.Now().UnixNano()

	for i := int64(0); i < n_sha256_trials; i++ {
		out = h.HashVals(buff.Bytes())
		buff.Reset()
		buff.Write(out)
	}

	d := float64(time.Now().UnixNano() - t1) / 1e9
	r := float64(n_sha256_trials) / d

	fmt.Printf("%s  Final hash: %x\n", prefix, buff.Bytes())
	fmt.Printf("%sRunning time: %g secs.\n", prefix, d)
	fmt.Printf("%s   Hash-rate: %g hashes-per-sec\n", prefix, r)
}
