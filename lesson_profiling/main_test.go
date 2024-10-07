package main

import "testing"

// go test -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out -x .

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fast()
	}
}

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slow()
	}
}

func BenchmarkOnStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OnStack()
	}
}
