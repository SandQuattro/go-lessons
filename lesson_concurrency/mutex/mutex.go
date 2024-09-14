package main

import (
	"fmt"
	"sync"
)

const (
	goroutinesNumber = 1000
	unlocked         = false
	locked           = true
)

type Mutex struct {
	state bool
}

// пытаемся взять мьютекс в цикле (spinlock)
// какие тут проблемы?
func (m *Mutex) Lock() {
	for m.state {
		// iteration by iteration...
		m.state = locked
	}
}

func (m *Mutex) Unlock() {
	m.state = unlocked
}

func main() {
	var mutex Mutex
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
