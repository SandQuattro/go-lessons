package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicMutex struct {
	state atomic.Bool
}

// Lock пытаемся взять мьютекс в цикле (spinlock)// какие тут проблемы?
func (m *AtomicMutex) Lock() {
	for m.state.Load() {
		// iteration by iteration...

		m.state.Store(locked)
	}
}

func (m *AtomicMutex) Unlock() {
	m.state.Store(unlocked)
}

func main() {
	var mutex AtomicMutex

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
