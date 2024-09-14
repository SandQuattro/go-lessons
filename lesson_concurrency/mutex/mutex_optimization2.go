package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicV2Mutex struct {
	state atomic.Bool
}

// Lock пытаемся взять мьютекс в цикле (spinlock)
func (m *AtomicV2Mutex) Lock() {
	for !m.state.CompareAndSwap(unlocked, locked) {
		// iteration by iteration...
	}

}

func (m *AtomicV2Mutex) Unlock() {
	m.state.Store(unlocked)
}

func main() {
	var mutex AtomicV2Mutex

	wg := sync.WaitGroup{}
	wg.Add(goroutinesNumber)

	value := 0
	for i := 0; i < goroutinesNumber; i++ {
		go func() {
			defer wg.Done()
			mutex.Lock()
			value++
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(value)
}
