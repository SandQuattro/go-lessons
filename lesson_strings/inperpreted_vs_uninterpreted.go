package main

import "fmt"

func main() {
	str := "\xF0\x9F\x98\x81"
	fmt.Println(str)

	raw := `\xF0\x9F\x98\x81`
	fmt.Println(raw)
}
