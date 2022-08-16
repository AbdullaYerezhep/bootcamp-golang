package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		io.Copy(os.Stdout, os.Stdin)
	} else {
		for i := range args {
			content, err := ioutil.ReadFile(args[i])
			if err != nil {
				printStr("ERROR: ")
				printStr(err.Error())
				z01.PrintRune('\n')
				printStr("exit status 1")
				z01.PrintRune('\n')
			}
			printStr(string(content))
		}
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
