package main

import "fmt"

const (
	PI = 3.14
)
var (
	a int
	b int
)

func main(){
	fmt.Println(PI)
	a,b := 1,2
	fmt.Println(a,b)
	fmt.Println("these are variables")
}
