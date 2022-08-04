package main

import "github.com/01-edu/z01"

func main() {
	PrintNbr(-123)
}

func PrintNbr(n int) {
	answer := ""

	for n > 10 || n < -10 {
		reminder := n % 10
		if reminder < 0 {
			reminder *= -1
		}
		answer += string(rune(reminder + 48))

		n /= 10
	}
	if n < 0 {
		z01.PrintRune('-')
		n*=-1
	}
	z01.PrintRune(rune(n+48))
	runes := []rune(answer)
	for i := len(answer)-1; i >= 0; i-- {
		z01.PrintRune(runes[i])
	}
}