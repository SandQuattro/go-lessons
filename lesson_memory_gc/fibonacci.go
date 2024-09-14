package main

import (
	"fmt"
)

func main() {
	fibonacci(0, 1)
}

func fibonacci(i, j int) {
	if i > 1000000000000000000 {
		return
	}
	fmt.Println(i)
	fibonacci(i+j, i)
}
