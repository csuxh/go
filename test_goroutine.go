package main

import 	(
    "fmt"
    "time"
)

func hello(){
    fmt.Println("from goroutine")
}

func main() {
    go hello()
    time.Sleep(1 * time.Second)
    fmt.Println("from Main")
}
