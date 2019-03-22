//defer in Go
package main

import (
	"fmt"
)

func worker() {
	defer del(10)
	defer del(20)
	fmt.Println("worker...")
}

func del(a int){
	fmt.Println("deleting ",a)
}

func main(){
	worker()
}
