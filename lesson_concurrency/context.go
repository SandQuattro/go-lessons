package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()

	stopped := make(chan struct{})
	go func() {
		// Используем буферизированный канал, как рекомендовано внутри signal.Notify функции
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// Блокируемся и ожидаем из канала quit - interrupt signal,
		// чтобы сделать gracefully shutdown с таймаутом в 10 сек
		<-quit

		fmt.Println("got termination signal")
		// Завершаем работу горутин
		// cancelFunc()

		// Получили SIGINT (0x2) или SIGTERM (0xf), выполняем graceful shutdown
		_, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		// здесь мы останавливаем сервер, закрываем ресурсы и тд...
		// if err := a.Srv.Shutdown(exitCtx); err != nil {
		// 	fmt.Println("gracefully shutdown error")
		// } else {
		// 	fmt.Println("Server stopped")
		// }

		close(stopped)
	}()

	fmt.Println("ok, now we are waiting for signal here...")
	<-stopped

	fmt.Println("bye, bye!")
}
