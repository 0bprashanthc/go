//even-ended number has same first and last digits
//how many even-ended numbers are there when two 4-digit numbers are multiplied

package main

import (
	"fmt"
)

func main(){
	count := 0

	for a:=1000; a<10000; a++ {
		for b:=a; b<10000; b++ {
			n := a*b
			s := fmt.Sprintf("%d",n)
			if s[0] == s[len(s)-1]{
				count++
			}
		}
	}
	fmt.Println("=== COUNT ===")
	fmt.Println(count)
}
