package main
import (
	"fmt"
	"log"
	"bufio"
	"os"
	"sync"
)

func readfile(filename string) {
	file,err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fmt.Println("======",filename,"======")
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	if err:=scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main(){
	files := []string{
		"files/1.txt",
		"files/2.txt",
		"files/3.txt",
		"files/4.txt",
		"files/5.txt",
		"files/6.txt",
		"files/7.txt",
		"files/8.txt",
		"files/9.txt",
		"files/10.txt",
	}
	var wg sync.WaitGroup
	for _,filename := range files{
		wg.Add(1)
		go func(filename string){
			readfile(filename)
			wg.Done()
		}(filename)
	}
	wg.Wait()
}
