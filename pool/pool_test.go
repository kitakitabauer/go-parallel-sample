package pool

import "testing"

func Benchmark_Pool1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool1("1")
	}
}

func Benchmark_Pool2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool2("2")
	}
}

func Benchmark_Pool3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool3("3")
	}
}
