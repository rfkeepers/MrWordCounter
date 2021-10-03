package main

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

const (
	minWordCount = 64
	maxWordCount = 256
)

var loremFields = strings.Fields(loremIpsum)
var rando *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
var result wordCounts

// builds n random strings containing some count of words between the min and max
// lengths specified.  Words are randomly chosen from the loremFields set.
func randomLoremBuilder(n, minLen, maxLen int) []string {
	resultSet := make([]string, n)

	for i := 0; i < n; i++ {
		input := []string{}
		length := rando.Intn(maxLen-minLen) + minLen

		for j := 0; j < length; j++ {
			input = append(input, loremFields[rando.Intn(len(loremFields))])
		}

		resultSet[i] = strings.Join(input, " ")
	}

	return resultSet
}

// assuming we don't need more than 1024 input strings, we can instantiate our
// test set with one call to this, and have each test run against different
// slices of lengths.  It's not perfect, but at least it should minimize any
// time spent recording a new random lorem creation.
var lorem1024 = randomLoremBuilder(1024, minWordCount, maxWordCount)
var lorem512 = lorem1024[:512]
var lorem256 = lorem1024[:256]
var lorem128 = lorem1024[:128]
var lorem64 = lorem1024[:64]
var lorem32 = lorem1024[:32]
var lorem8 = lorem1024[:8]
var lorem2 = lorem1024[:2]

// channeled

func benchmarkChanneled(goRoutineCount int, inputs []string, b *testing.B) {
	var r wordCounts
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		r = mrChanneledCounter(goRoutineCount, inputs)
	}
	// something something prevent compiler optimization
	result = r
}

// one routine
func BenchmarkChanneled_1_1024(b *testing.B) {
	benchmarkChanneled(1, lorem1024, b)
}
func BenchmarkChanneled_1_512(b *testing.B) {
	benchmarkChanneled(1, lorem512, b)
}
func BenchmarkChanneled_1_256(b *testing.B) {
	benchmarkChanneled(1, lorem256, b)
}
func BenchmarkChanneled_1_128(b *testing.B) {
	benchmarkChanneled(1, lorem128, b)
}
func BenchmarkChanneled_1_64(b *testing.B) {
	benchmarkChanneled(1, lorem64, b)
}
func BenchmarkChanneled_1_32(b *testing.B) {
	benchmarkChanneled(1, lorem32, b)
}
func BenchmarkChanneled_1_8(b *testing.B) {
	benchmarkChanneled(1, lorem8, b)
}
func BenchmarkChanneled_1_2(b *testing.B) {
	benchmarkChanneled(1, lorem2, b)
}

// 256 routines
func BenchmarkChanneled_256_1024(b *testing.B) {
	benchmarkChanneled(256, lorem1024, b)
}
func BenchmarkChanneled_256_512(b *testing.B) {
	benchmarkChanneled(256, lorem512, b)
}
func BenchmarkChanneled_256_256(b *testing.B) {
	benchmarkChanneled(256, lorem256, b)
}
func BenchmarkChanneled_256_128(b *testing.B) {
	benchmarkChanneled(256, lorem128, b)
}
func BenchmarkChanneled_256_64(b *testing.B) {
	benchmarkChanneled(256, lorem64, b)
}
func BenchmarkChanneled_256_32(b *testing.B) {
	benchmarkChanneled(256, lorem32, b)
}
func BenchmarkChanneled_256_8(b *testing.B) {
	benchmarkChanneled(256, lorem8, b)
}
func BenchmarkChanneled_256_2(b *testing.B) {
	benchmarkChanneled(256, lorem2, b)
}

// sliced

func benchmarkSliced(goRoutineCount int, inputs []string, b *testing.B) {
	var r wordCounts
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		r = mrSlicedCounter(goRoutineCount, inputs)
	}
	// something something prevent compiler optimization
	result = r
}

// one routine
func BenchmarkSliced_1_1024(b *testing.B) {
	benchmarkSliced(1, lorem1024, b)
}
func BenchmarkSliced_1_512(b *testing.B) {
	benchmarkSliced(1, lorem512, b)
}
func BenchmarkSliced_1_256(b *testing.B) {
	benchmarkSliced(1, lorem256, b)
}
func BenchmarkSliced_1_128(b *testing.B) {
	benchmarkSliced(1, lorem128, b)
}
func BenchmarkSliced_1_64(b *testing.B) {
	benchmarkSliced(1, lorem64, b)
}
func BenchmarkSliced_1_32(b *testing.B) {
	benchmarkSliced(1, lorem32, b)
}
func BenchmarkSliced_1_8(b *testing.B) {
	benchmarkSliced(1, lorem8, b)
}
func BenchmarkSliced_1_2(b *testing.B) {
	benchmarkSliced(1, lorem2, b)
}

// 256 routines
func BenchmarkSliced_256_1024(b *testing.B) {
	benchmarkSliced(256, lorem1024, b)
}
func BenchmarkSliced_256_512(b *testing.B) {
	benchmarkSliced(256, lorem512, b)
}
func BenchmarkSliced_256_256(b *testing.B) {
	benchmarkSliced(256, lorem256, b)
}
func BenchmarkSliced_256_128(b *testing.B) {
	benchmarkSliced(256, lorem128, b)
}
func BenchmarkSliced_256_64(b *testing.B) {
	benchmarkSliced(256, lorem64, b)
}
func BenchmarkSliced_256_32(b *testing.B) {
	benchmarkSliced(256, lorem32, b)
}
func BenchmarkSliced_256_8(b *testing.B) {
	benchmarkSliced(256, lorem8, b)
}
func BenchmarkSliced_256_2(b *testing.B) {
	benchmarkSliced(256, lorem2, b)
}
