package utils

import (
	"runtime"
	"testing"
)

func Benchmark_X(b *testing.B) {
	runtime.GOMAXPROCS(8)
	b.SetParallelism(4000)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			X()
		}
	})
}
