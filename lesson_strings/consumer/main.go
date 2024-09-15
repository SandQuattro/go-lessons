package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"runtime"
)

func main() {
	listener, err := net.Listen("tcp4", ":6676")
	if err != nil {
		return
	}
	defer listener.Close()

	buf := make([]byte, 10)

	fmt.Println("Listening on :6676")

	// смотрим в рантайме, сколько программа использует памяти
	var memStat runtime.MemStats
	runtime.ReadMemStats(&memStat)

	for {
		socket, err := listener.Accept()
		if err != nil {
			return
		}

		str := make([]byte, 0)
		for {
			numBytes, err := socket.Read(buf)
			if errors.Is(err, io.EOF) {
				fmt.Printf("[INCOMING MESSAGE] %s\n", string(str))
				fmt.Println("!!! Done reading from client")
				break
			}

			if err != nil {
				panic(err)
			}

			str = append(str, buf[:numBytes]...)

			fmt.Printf("read %d bytes from client, data:%s\n", numBytes, string(buf[:numBytes]))

			if buf[numBytes-1] == 9 {
				fmt.Printf("[INCOMING MESSAGE] %s\n", string(str))
				str = str[len(str):]

				// смотрим в рантайме, сколько программа использует памяти
				var memStatNow runtime.MemStats
				runtime.ReadMemStats(&memStatNow)

				memConsumed := float64(memStatNow.Sys-memStat.Sys) / 1024 / 1024
				runtime.GC()
				fmt.Printf("Memory consumed %f MB\n", memConsumed)
			}
		}
	}

}
