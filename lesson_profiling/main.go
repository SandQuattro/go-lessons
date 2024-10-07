package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

type MyStruct struct {
	data [1 << 20]byte // 1 MB данных
}

// 1) go tool pprof lesson_profiling/main.go http://127.0.0.1:8080/debug/pprof/profile
// 1) curl -o cpu.pprof -location 'http://localhost:8080/debug/pprof/profile?seconds=10' CPU ИЛИ
// 1) curl -o heap.pprof -location 'http://localhost:8080/debug/pprof/heap?seconds=10' HEAP ИЛИ
// 1) curl -o go.pprof -location 'http://127.0.0.1:8080/debug/pprof/goroutine?seconds=10' GOROUTINE
// 2) bombardier -c 125 -d 1ms -n 100 http://localhost:8080
// 3) go tool pprof <pprof file>
// 4) top
// 5) Если на первом месте у нас runtime.kevent, то это net poller для macOS
// It's the network poller. There are multiple implementations:
//
// epoll for linux
// kevent queue for darwin and a bunch of others. Basically, this is time spent waiting for I/O
// 6) выполняем команду list runtime.kevent: и видим вызов C функции kevent_trampoline:
// ret := libcCall(unsafe.Pointer(abi.FuncPCABI0(kevent_trampoline)), unsafe.Pointer(&kq))
// 7) web

// Открываем профили в браузере
// 8) go tool pprof -http=:7272 <pprof file>
// 9) go tool pprof -http=:7272 -diff_base <pprof file old> <pprof file new>
// 10) go tool pprof lesson_profiling/main.go http://127.0.0.1:8080/debug/pprof/goroutine
// 10) go tool pprof -http=:7272 <pprof file>
func main() {
	server := &http.Server{
		ReadHeaderTimeout: 5 * time.Second,
		Addr:              ":8080",
	}

	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
		<-signals

		shutCtx, shutFunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutFunc()

		_ = server.Shutdown(shutCtx)
		cancelFunc()
	}()

	go func() {
		http.HandleFunc("/", handleRoot)
		_ = server.ListenAndServe()
	}()

	<-ctx.Done()
	fmt.Println("Server is shutting down... bye, bye")
}

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	//for i := 0; i < 100; i++ {
	//	s := new(MyStruct)
	//	for j := 0; j < len(s.data); j++ {
	//		s.data[j] = byte(i)
	//	}
	//	fmt.Printf("Allocated %d MB\n", i+1)
	//}

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
		runtime.Gosched()
	}

	wg.Wait()

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello"))
}

func Fast() int {
	acc := new(int)
	for i := 0; i < 10; i++ {
		acc2 := new(int)
		*acc2 = *acc + 1
		acc = acc2
	}
	return *acc
}

func Slow() int {
	acc := new(int)
	for i := 0; i < 1000; i++ {
		acc2 := new(int)
		*acc2 = *acc + 1
		acc = acc2
	}
	return *acc
}

func OnStack() {
	for i := 0; i < 1000; i++ {
		a := 100
		_ = a
	}
}
