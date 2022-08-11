package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for _, w := range arguments[1:] {
		for j := range w {
			z01.PrintRune(rune(w[j]))
		}
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')
}
