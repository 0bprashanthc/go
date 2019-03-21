//if statement
package main

import (
	"fmt"
)

func main() {
	x,y := 10,20

	if x > 5 {
		fmt.Println("x > 5 ")
	} else {
		fmt.Println("x < 5 ")
	}
	if frac := x/y; frac > 2 {
		fmt.Println("true")
	}else {
		fmt.Println("false")
	}
}
