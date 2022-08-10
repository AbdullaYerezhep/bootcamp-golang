package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programname := os.Args[1]
	for i := 0; i < len(os.Args[0]); i++ {
		z01.PrintRune(rune(programname[i]))
	}
}
