package main
import (
	"fmt"
	"sync"
	"time"
)

func do_something(some string, i int){
	fmt.Println(some,i)
	fmt.Println(i)
}

func main() {
	var wg sync.WaitGroup
	for i:=0; i<10; i++ {
		wg.Add(1)
		go do_something("something_",i)
		wg.Done()
	}
	time.Sleep(10*time.Millisecond)
	wg.Wait()
}
