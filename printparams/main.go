package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	for i := range arguments {
		for _, char := range arguments[i] {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
