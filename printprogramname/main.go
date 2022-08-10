package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programname := os.Args[0][2:]
	for i := range programname {
		z01.PrintRune(rune(programname[i]))
	}
	z01.PrintRune('\n')
}
