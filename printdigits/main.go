package main

import "github.com/01-edu/z01"

func main() {
	numbers := "0123456789"
	chars := []rune(numbers)
	for i := 0; i < len(chars); i++ {
		letter := chars[i]
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
