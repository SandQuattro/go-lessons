package sharded_counter

import (
	"runtime"
	"sync"
	"testing"
)

// original case, just mutex
func BenchmarkMutexCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	counter := MutexCounter{}

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			defer wg.Done()
			// bench
			for j := 0; j < b.N; j++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
}

// optimization step 1, use RW Mutex
func BenchmarkRWMutexCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	counter := RWMutexCounter{}

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			defer wg.Done()
			// bench
			for j := 0; j < b.N; j++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
}

// optimization step 2, use atomic counter
func BenchmarkAtomicCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	counter := AtomicCounter{}

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			defer wg.Done()
			// bench
			for j := 0; j < b.N; j++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
}

// optimization step 3, use sharded atomic counter
func BenchmarkShardedAtomicCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	counter := ShardedAtomicCounter{}

	for i := 0; i < runtime.NumCPU(); i++ {
		i := i
		go func() {
			defer wg.Done()
			// bench
			for j := 0; j < b.N; j++ {
				counter.Increment(i)
			}
		}()
	}
	wg.Wait()
}

// final optimization step 4, use aligned sharded atomic counter
func BenchmarkAlignedShardedCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	counter := AlignedShardedAtomicCounter{}

	for i := 0; i < runtime.NumCPU(); i++ {
		i := i
		go func() {
			defer wg.Done()
			// bench
			for j := 0; j < b.N; j++ {
				counter.Increment(i)
			}
		}()
	}
	wg.Wait()
}
