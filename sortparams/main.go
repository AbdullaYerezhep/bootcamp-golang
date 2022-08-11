package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	for i := len(arguments) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arguments[j] > arguments[j+1] {
				arguments[j], arguments[j+1] = arguments[j+1], arguments[j]
			}
		}
	}

	for i := range arguments {
		for _, char := range arguments[i] {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
