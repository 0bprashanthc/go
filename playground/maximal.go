//calculate maximum value in a slice

package main

import (
	"fmt"
)

func main() {
	nums := []int{3004,2,6,1,8,3}
	fmt.Println(nums)
	var max int
	fmt.Println(max)
	for _,value := range nums {
		if value > max{
			max = value
		}
	}
	fmt.Println("=== maximum value is ===")
	fmt.Println(max)
}
