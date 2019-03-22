//pointers in Go. Simple Example

package main

import (
	"fmt"
)

func appendd(s *string){
	*s = *s + *s
}

func main() {
	s := " I am Prashanth "
	appendd(&s)
	fmt.Println(s)
}
