package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//this program acts as a package manager to go

func main() {
	if len(os.Args) > 1 {
		url := os.Args[1]
		client := &http.Client{}
		resp, err := client.Get(url)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		f, err := os.Create("/Users/kurian/Downloads/hello.zip")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.Write(body)
		fmt.Printf("%s", url)
	}
	return
}
