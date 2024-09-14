package main

import "fmt"

func main() {
    ch := generator()

    for i := 0; i < 5; i++ {
        value := <-ch
        fmt.Println("Value:", value)
    }
}

func generator()