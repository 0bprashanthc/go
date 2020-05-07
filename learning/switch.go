package main

import (
    "fmt"
    "runtime"
)

func main(){
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("mac OS")
    case "linux":
        fmt.Println("linux")
    }
}
