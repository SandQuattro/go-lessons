package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 1) curl -location 'http://localhost:8080/debug/pprof/profile?seconds=10' CPU ИЛИ
// 1) curl -location 'http://localhost:8080/debug/pprof/heap?seconds=5'
// 2) bombardier -c 125 -d 1ms -n 100 http://localhost:8080
// 3) go tool pprof <pprof file>
// 4) top
// 5) Если на первом месте у нас runtime.kevent, то это net poller для mac os
// It's the network poller. There are multiple implementations:
//
// epoll for linux
// kevent queue for darwin and a bunch of others. Basically, this is time spent waiting for I/O
// 6) выполняем команду list runtime.kevent: и видим вызов C функции kevent_trampoline:
// ret := libcCall(unsafe.Pointer(abi.FuncPCABI0(kevent_trampoline)), unsafe.Pointer(&kq))
// 7) web
// 8) go tool pprof -http=:7272 <pprof file>
// 9) go tool pprof -http=:7272 -diff_base <pprof file old> <pprof file new>
func main() {
	server := &http.Server{
		Addr: ":8080",
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

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}
