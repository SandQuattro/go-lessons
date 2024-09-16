package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

//go:embed data
var sentence string

func main() {

	conn, err := net.Dial("tcp", ":6676")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	lines := strings.Split(sentence, "\n")

	for {
		randomLine := lines[rand.Intn(len(lines))] + "\n"

		fmt.Println(randomLine)
		fmt.Println()
		conn.Write([]byte(randomLine))
		i := rand.Intn(20) % 10
		time.Sleep(time.Duration(i) * time.Second)

	}

}
