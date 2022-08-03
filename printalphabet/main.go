package main

import "github.com/01-edu/z01"

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	letters := []rune(alphabet)
	for i := 0; i < len(letters); i++ {
		letter := letters[i]
		z01.PrintRune(letter)

	}
}
