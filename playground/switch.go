//switch statement in go
package main

import (
	"fmt"
)

func main(){
	a := 10
	switch a {
		 case 1:
			 fmt.Println("x is 1")
		 case 10:
			 fmt.Println("x is 10")
		 case 100:
			 fmt.Println("x is 100")
		 default:
			 fmt.Println("x is default")
	}
	switch {
	case a > 10:
		fmt.Println("a > 10")
	default:
		fmt.Println("default")
	}
}

