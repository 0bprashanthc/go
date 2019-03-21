//maps/HashMaps in Go
package main

import (
	"fmt"
)

func main(){
	names := map[string]int{}
	names["a"] = 123
	names["b"] = 123
	names["c"] = 123
	names["d"] = 123
	names["e"] = 123
	names["f"] = 123
	names["g"] = 123
	names["h"] = 123
	names["i"] = 123
	names["g"] = 13
	fmt.Println(names)
	i := 0
	for i<len(names) {
		value, found := names["a"]
		fmt.Println(value, found)
		fmt.Println("a : ", value)
		break
	}
	delete(names, "a")
	for key := range names {
		fmt.Println(key, names[key])
	}
	for key,value := range names{
		fmt.Println(key," : ", value)
	}
}
