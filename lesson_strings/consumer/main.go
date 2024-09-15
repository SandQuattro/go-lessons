package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"net"
	"net/http"
)

type Server struct {
	upgrader *websocket.Upgrader
	ch       chan string
}

func main() {
	server := Server{
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		ch: make(chan string),
	}

	go func() {
		listener, err := net.Listen("tcp4", ":6676")
		if err != nil {
			return
		}
		defer listener.Close()

		buf := make([]byte, 10)

		fmt.Println("Listening on :6676")
		for {
			socket, err := listener.Accept()
			if err != nil {
				return
			}

			str := make([]byte, 0)
			for {
				numBytes, err := socket.Read(buf)
				if errors.Is(err, io.EOF) {
					server.ch <- fmt.Sprintf("[INCOMING MESSAGE] %s\n", string(str))
					fmt.Println("!!! Done reading from client")
					break
				}

				if err != nil {
					panic(err)
				}

				str = append(str, buf[:numBytes]...)

				fmt.Printf("read %d bytes from client, data:%s\n", numBytes, string(buf[:numBytes]))

				if buf[numBytes-1] == 9 {
					server.ch <- fmt.Sprintf("[INCOMING MESSAGE] %s\n", string(str))
					str = str[len(str):]
				}
			}
		}
	}()

	http.HandleFunc("/ws", server.handleWs)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}

func (s *Server) handleWs(w http.ResponseWriter, r *http.Request) {
	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	for {
		go func() {
			for {
				_, msg, err := ws.ReadMessage()
				if err != nil {
					fmt.Println("Error reading message:", err)
					break
				}
				fmt.Printf("Received message: %s\n", msg)
			}
		}()

		err = ws.WriteMessage(websocket.TextMessage, []byte(<-s.ch))
		if err != nil {
			return
		}
	}
}
