package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	requesturl := "http://www.google.com"
	response,err := http.Get(requesturl)
	if err != nil {
		fmt.Println("err: ",err)
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		r, err := ioutil.ReadAll(response.Body)
		if  err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(r))
	}
}