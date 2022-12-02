package main

import (
    "log"
)

func main(){
    log.Println("1 + 1 is ", Add(1,1))
}

func Add(x int, y int) int {
    return x + y
}
