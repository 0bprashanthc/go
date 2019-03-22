/*
function that gets a URL and returns the
value of Content-Type response HTTP header.
    - func should return an error if it can't 
      perform a GET request
*/
package main

import (
	"fmt"
	"net/http"
)

func getContentTypeHTTP(url string) (string, error){
	var contentType string
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil{
		return "nil", err
	} else {
                contentType = resp.Header.Get("Content-Type")
                if contentType != "" {
                        return contentType, nil
                }else {
                        return "",fmt.Errorf("can't find content-type")
                }
	}
}

func main(){
	url := "https://linkedin.com"
	contentType, err := getContentTypeHTTP(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(contentType)
	}
}
