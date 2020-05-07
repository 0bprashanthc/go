package main

import "fmt"

func sqrt(x float32) (z float32){
    var tmp float32
    z = float32(1)
    count := 0
    for count < 10{
        z -= (z*z-x)/(2*z)
        if tmp == z {
            break
        }
        tmp = z
        count++
    }
    return
}

func main(){
    fmt.Println("sqrt(float32(1600)):",sqrt(float32(1600)))
}
