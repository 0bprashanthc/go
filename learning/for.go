package main

import "fmt"

func main(){
    for i:=0; i<10; i++ {
        fmt.Println("i: ",i)
    }
    sum := 0
    for ; sum <25; {
        if sum%2 == 0{
            fmt.Println(sum)
        }
        sum++
    }
    //while in go is 'for' itself
    sum = 0
    for sum<25{
        if sum%2 == 0 {
            fmt.Println(sum)
        }
        sum++
    }
    for {
    }
}
