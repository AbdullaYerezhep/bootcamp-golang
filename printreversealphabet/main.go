package main

import "github.com/01-edu/z01"

func main() {
	alphabet := "zyxwvutsrqponmlkjihgfedcba"
	letters := []rune(alphabet)
	for i := 0; i < len(letters); i++ {
		letter := letters[i]
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
