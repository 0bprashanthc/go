//count number of words from text
package main

import (
	"fmt"
	"strings"
)

func main(){
	text := `
	i am prashanth
	i am bangalore
	i work at vmware
	i like music
	`
	fmt.Println(text)
	words := strings.Fields(text)
	counts := map[string]int{}
	for _,word := range words{
		counts[strings.ToLower(word)]++
	}
	for key,word := range counts {
		fmt.Println(key, word)
	}
}
