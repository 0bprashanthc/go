package main
import (
	"net/http"
	"fmt"
	"sync"
)

func getContentType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error")
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Println(url)
	fmt.Println(ctype)
}

func main() {
	urls := []string{
		"https://linkedin.com",
		"https://google.com",
		"https://facebook.com",
		"https://vmware.com",
		}
	var wg sync.WaitGroup
	for _,url := range urls{
		wg.Add(1)
		go func(url string){
			getContentType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

