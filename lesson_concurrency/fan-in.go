package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			ch1 <- num
		}
		close(ch1)
	}()

	go func() {
		for _, num := range []int{10, 20, 30} {
			ch2 <- num
		}
		close(ch2)
	}()

	go func() {
		for _, num := range []int{100, 200, 300} {
			ch3 <- num
		}
		close(ch3)
	}()

	for val := range merge(ch1, ch2, ch3) {
		fmt.Println(val)
	}
}

func merge[T any](chans ...chan T) chan T {
	return result
}
