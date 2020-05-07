package main

import "fmt"

func main(){
    fmt.Println("add(10,11): ", add(10,11))
    a,b := 10,11
    a,b = swap(a,b)
    fmt.Println("a:",a,"b:",b,"swap(a,b): ",a,b)
}

func add(x, y int) int {
    return x+y
}

func swap(x, y int) (int, int) {
    return y,x
}
