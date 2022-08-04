package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	chars := []rune(s)
	for i := 0; i < len(chars); i++ {
		z01.PrintRune(chars[i])
	}
}
