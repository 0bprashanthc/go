//stderr from fmt package in Go
package main

import (
	"fmt"
)

func divide(a int, b int) (int, error){
	if b == 0 {
		return 0, fmt.Errorf("b cannot be 0")
	} else {
		return a/b, nil
	}
}

func main(){
	result, err := divide(10,10)
	fmt.Println(result)
	fmt.Println(err)
}
