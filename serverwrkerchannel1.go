package main

import (  
    "fmt"
)


func main() {  
    ch := make(chan string, 2)
    ch <- "sudha"
    ch <- "nayuni"
    fmt.Println(<- ch)
    fmt.Println(<- ch)
}