//slices in Go
package main

import (
	"fmt"
)

func main(){
	fruits := []string{"apple","mango","banana"}
	fmt.Println(len(fruits))
	fmt.Println(fruits[1])
	fmt.Println(fruits[1:])

	fmt.Println("=== index ===")
	for i:=0; i<len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	fmt.Println("=== index, value@index ===")
	for i := range fruits {
		fmt.Println(i, fruits[i])
	}

	fmt.Println("=== index, value ===")
	for i,name := range fruits {
		fmt.Println(i, name)
	}

	fmt.Println("=== _, value ===")
	for _,name := range fruits {
		fmt.Println(name)
	}

	fmt.Println("==> APPENDING...")
	fruits = append(fruits, "guava")
	fmt.Println(fruits)
}
