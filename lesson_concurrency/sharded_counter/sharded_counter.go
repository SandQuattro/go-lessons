package sharded_counter

import (
	"sync"
	"sync/atomic"
)

// RWMutexCounter case 1
type RWMutexCounter struct {
	value int
	mx    sync.RWMutex
}

func (c *RWMutexCounter) Increment() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.value++
}

// MutexCounter case 2
type MutexCounter struct {
	value int
	mx    sync.Mutex
}

func (c *MutexCounter) Increment() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.value++
}

// AtomicCounter case 3
type AtomicCounter struct {
	value atomic.Int32
}

func (c *AtomicCounter) Get() int32 {
	return c.value.Load()
}

func (c *AtomicCounter) Increment() {
	c.value.Add(1)
}

// ShardedAtomicCounter для 3 этапа оптимизации
type ShardedAtomicCounter struct {
	shards [14]AtomicCounter
}

// AlignedAtomicCounter case 4
type AlignedAtomicCounter struct {
	value     atomic.Int32
	alignment [60]byte
}

// AlignedShardedAtomicCounter case 4
type AlignedShardedAtomicCounter struct {
	shards [14]AlignedAtomicCounter
}

func (c *ShardedAtomicCounter) ShardedGet(idx int) int32 {
	var value int32
	for i := 0; i < 10; i++ {
		value += c.shards[idx].Get()
	}
	return value
}

func (c *ShardedAtomicCounter) Increment(idx int) {
	c.shards[idx].value.Add(1)
}

func (c *AlignedShardedAtomicCounter) Increment(idx int) {
	c.shards[idx].value.Add(1)
}
