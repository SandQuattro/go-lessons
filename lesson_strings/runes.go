package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// иммутабельность строк
	str := "test"
	fmt.Printf("string pointer %p, value: %v\n", &str, str)

	b := []byte(str)
	fmt.Printf("pointer %p, value: %v\n", b, b)

	b[1] = 'o'
	fmt.Printf("pointer %p, value: %v\n", b, b)
	fmt.Printf("string pointer %p, value: %v\n", &str, str)

	str = string(b)
	fmt.Printf("pointer %p, value: %v\n", &str, str)

	str = "П😁"
	fmt.Printf("pointer %p, type:%T, value: %v, len: %d\n", &str, str, str, len(str))
	fmt.Printf("string actual lenght: %d, %d\n", utf8.RuneCountInString(str), len([]rune(str)))

	b = []byte(str)
	fmt.Printf("pointer %p, binary:%b value: %v \n", b, b, b)
}
