package main

import (
    "fmt"
    "time"
)

func write(ch chan int) {
    for i :=0; i < 5; i++ {
        ch <- i
        fmt.Println("Successfully write", i, "to chan")
    }
    close(ch)
}

func main() {
    ch := make(chan int, 2)
    go write(ch)
    time.Sleep(3 * time.Second)
    for v := range ch {
        fmt.Println("read ", v, )
        time.Sleep(3 * time.Second)
    }
 }
