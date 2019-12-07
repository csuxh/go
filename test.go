package main

import "fmt"

func main() {
    //var width, height int = 100, 50 // 声明多个变量
    //fmt.Println("width is", width, "height is", heigh)
    // var mapTest map[string]int
    mapTest := make(map[string]int)
    mapTest["jack"] = 30
    mapTest["Alex"] = 25 
    fmt.Println(mapTest)
    for key, value := range mapTest{
        fmt.Println(key, value)  
    }
    // del(mapTest, key), len(mapTest), 
}
