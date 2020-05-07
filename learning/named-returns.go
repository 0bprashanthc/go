package main

import "fmt"

func main(){
    a,b := naked_return(10)
    fmt.Println("naked_return(10): ",a,b)
}

func naked_return(sum int) (x,y int){
    x = sum*4
    y = x/3
    return
}
