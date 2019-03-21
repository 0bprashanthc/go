//strings in go
package main

import (
	"fmt"
)

func main() {
	s := "I am Prashanth!"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(s[0])
	fmt.Println("I "+s[2:])
	s = `
	i am from bangalore
	India.
	`
	fmt.Println(s)
}
