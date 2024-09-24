package main

import (
	"math/rand"
	"sync"
	"time"
)

var data = []string{"cat", "dog", "squirrel", "bear", "mouse"}
var cond = sync.NewCond(&sync.Mutex{})
var pokemon = ""

func main() {
	// Consumer
	go func() {
		cond.L.Lock()
		defer cond.L.Unlock()

		// waits until Pikachu appears
		for pokemon != "squirrel" {
			cond.Wait()
		}
		println("Caught " + pokemon)
		pokemon = ""
	}()

	// Producer
	go func() {
		// Every 1ms, a random Pokémon appears
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond)

			cond.L.Lock()
			pokemon = data[rand.Intn(len(data))]
			cond.L.Unlock()

			cond.Signal()
		}
	}()

	time.Sleep(100 * time.Millisecond) // lazy wait
}
