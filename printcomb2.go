package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			z01.PrintRune(rune(i/10 + 48))
			z01.PrintRune(rune(i%10 + 48))
			z01.PrintRune(32)
			z01.PrintRune(rune(j/10 + 48))
			z01.PrintRune(rune(j%10 + 48))
			if(i != 98) {
				z01.PrintRune(44)
				z01.PrintRune(32)
			}
		}
	}
	z01.PrintRune('\n')
}
