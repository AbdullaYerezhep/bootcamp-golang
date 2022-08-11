package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	isupper := false

	if args[0] == "--upper" {
		isupper = true
		args = args[1:]
	}
	for i := range args {

		if isupper && rune(StringToInt(args[i])+64) >= 'A' && rune(StringToInt(args[i])+64) <= 'Z' {
			args[i] = string(rune(StringToInt(args[i]) + 64))
		} else if !isupper && rune(StringToInt(args[i])+96) >= 'a' && rune(StringToInt(args[i])+96) <= 'z' {
			args[i] = string(rune(StringToInt(args[i]) + 96))
		} else {
			args[i] = string(' ')
		}

		for _, char := range args[i] {
			z01.PrintRune(char)
		}

	}
	z01.PrintRune('\n')
}

func StringToInt(s string) int {
	runes := []rune(s)

	var arr []int

	n := 0
	for i := range s {
		if runes[i] <= '9' && runes[i] >= '0' {
			arr = append(arr, int(runes[i]-48))
		}
	}

	for i := 0; i < len(arr); i++ {
		n = n*10 + arr[i]
	}

	return n
}
