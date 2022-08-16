package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("File name missing")
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
	} else {
		content, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s", content)
	}
}
