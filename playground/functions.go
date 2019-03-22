package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a+b
}

func addAndsubtract(a int, b int) (int, int) {
	return a+b, a-b
}

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) {
	n *= 2
}

func doublePtr(n *int){
	*n = *n * 2
}

func main() {
	a,b := 10,20
	ad := add(a,b)
	a,s := addAndsubtract(a,b)
	fmt.Println(ad)
	fmt.Println(a,s)
	vals := []int{1,2,3,3,4}
	fmt.Println(vals)
	doubleAt(vals, 3)
	fmt.Println(vals)
	val := 12
	fmt.Println(val)
	double(val)
	fmt.Println(val)
	doublePtr(&val)
	fmt.Println(val)
}
