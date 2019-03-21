/*
	iterate from 1-15
	if number divisible by 3: fizz
	if number divisible by 5: buzz
	if number divisible by 3 and 5: fizzbuzz
*/
package main

import (
	"fmt"
)

func main(){
	a := 1
	for a < 16 {
		switch {
		case a%3==0 && a%5==0:
			fmt.Println("fizzbuzz")
		case a%3==0:
			fmt.Println("fizz")
		case a%5==0:
			fmt.Println("buzz")
		default:
			fmt.Println(a)
		}
		a++
	}
}
