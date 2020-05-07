package main

import (
    "fmt"
    "math"
)

func main(){
    if i := math.Sqrt(1); i < 2 {
        fmt.Println(i)
    } else {
        fmt.Println("something else")
    }
}
