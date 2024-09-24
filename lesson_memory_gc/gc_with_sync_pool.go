package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/trace"
	"sync"
)

type MyPoolStruct struct {
	data [1 << 20]byte // 1 MB данных
}

// go tool trace trace.out
func main() {
	// отслеживаем изменение в куче
	debug.SetGCPercent(-1)
	//debug.SetGCPercent(10)
	//debug.SetGCPercent(100)
	//debug.SetGCPercent(1000)

	// Запись в trace файл
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()

	// смотрим в рантайме, сколько программа использует памяти
	var memStat runtime.MemStats
	runtime.ReadMemStats(&memStat)

	pool := &sync.Pool{
		New: func() interface{} {
			return new(MyPoolStruct)
		},
	}

	for i := 0; i < 100; i++ {
		s := pool.Get().(*MyPoolStruct) // new(MyStruct)
		for j := 0; j < len(s.data); j++ {
			s.data[j] = byte(i)
		}

		fmt.Printf("Allocated new %d MB struct\n", i+1)

		pool.Put(s)
		// if i == 50 {
		// 	// Вызов сборщика мусора вручную
		// 	runtime.GC()
		// 	fmt.Println("Garbage collector invoked")
		// }
	}

	// смотрим в рантайме, сколько программа использует памяти
	var memStatNow runtime.MemStats
	runtime.ReadMemStats(&memStatNow)

	memConsumed := float64(memStatNow.Sys-memStat.Sys) / 1024 / 1024
	runtime.GC()
	fmt.Printf("Memory consumed %f MB\n", memConsumed)

}
