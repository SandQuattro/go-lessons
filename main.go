package main

import "fmt"

func main() {
	ch := make(chan int, 0)
	//ch := make(chan int)
	//go func() {
	//	<-ch
	//}()

	ch <- 1
  fmt.Println(<- ch)
}
