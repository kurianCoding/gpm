package main

import (
	"fmt"
	"os"
)

//this program acts as a package manager to go

func main() {
	if len(os.Args) > 1 {
		url := os.Args[1]
		fmt.Printf("%s", url)
	}
	return
}
