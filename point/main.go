package main

import (
	"github.com/01-edu/z01"
)

func setPoint(ptr []int) []int {
	x := 0
	y := 1
	ptr[x] = 42
	ptr[y] = 21
	return ptr
}

func main() {
	points := []int{0, 1}

	points = setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune('4')
	z01.PrintRune('2')
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune('2')
	z01.PrintRune('1')
	z01.PrintRune('\n')
}
