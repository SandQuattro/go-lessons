package main

import "testing"

// go test -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out -x .

// cpu: Apple M3 Max

// Тест и кол-во ядер   кол-во выполнений    скорость выполнения   память на операцию  аллокаций на операцию
// BenchmarkFast-14        14348666                78.99 ns/op           80 B/op         10 allocs/op
// BenchmarkSlow-14          171397              7046 ns/op            8000 B/op       1000 allocs/op
// BenchmarkOnStack-14      4665962               256.9 ns/op             0 B/op          0 allocs/op

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
