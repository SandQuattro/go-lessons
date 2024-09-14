package main

import (
	"fmt"
)

func f() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
	fmt.Printf("the value is %v.\n", data)
}

func main() {
	f()
}
