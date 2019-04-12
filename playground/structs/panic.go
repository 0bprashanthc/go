package main

import (
	"os"
)

func main(){
	_, err := os.Open("txt")
	strconv.Atoi("10")
	if err != nil {
		panic(err)
	}
}
