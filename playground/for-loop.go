//for loops in go
package main

import (
	"fmt"
)

func main(){
	for i:=0; i <3; i++ {
		if i > 10 {
			break
		}
		if i < 0 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("=====")
	a := 0
	for a < 3{
		fmt.Println(a)
		a++
	}
	fmt.Println("=====")
	b := 0
	for {
		fmt.Println(b)
		if b > 2 {
			fmt.Println("b > 2; so breaking...")
			break
		}
		b++
	}
}
