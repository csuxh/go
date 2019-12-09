package main

import (
    "fmt"
)

func hello(flag chan bool) {
    fmt.Println("Hello from function")
    flag <- true
}

func main() {
    flag := make(chan bool)
    go hello(flag)
    <-flag
    fmt.Println("main process")
}
