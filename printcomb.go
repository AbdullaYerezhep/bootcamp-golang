package piscine

import "github.com/01-edu/z01"


func PrintComb() {
	for i := '0'; i < '9'; i++ {
		for j := i + 1; j < '9'; j++ {
			for x := j + 1; x < '9'; x++ {
				if i < 7 {
					z01.PrintRune(i)
					if j < 8 {
						z01.PrintRune(j)
						z01.PrintRune(x)
					}
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
