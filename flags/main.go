package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	insert := false
	order := false
	o := false
	in := false
	var order_elem string
	var insert_str string
	n := 0

	if len(arguments) == 0 {
		error()
		n += 1
	} else if len(arguments[0]) >= 2 && arguments[0][:2] == "-h" {
		error()
		n += 1
	} else if len(arguments[0]) >= 3 && arguments[0][:3] == "--h" {
		error()
		n += 1
	}

	if n == 0 {
		for i := range arguments { // est' li ins,or
			if len(arguments[i]) >= 7 && arguments[i][:2] == "--" || len(arguments[i]) >= 2 && arguments[i][0] == '-' {
				if arguments[i][:2] == "--" && arguments[i][2:7] == "order" {
					order = true
					order_elem = arguments[i+1]
				} else if arguments[i][:2] == "-o" {
					o = true
					order_elem = arguments[i+1]
				} else if arguments[i][:2] == "--" && arguments[i][2:8] == "insert" {
					insert = true
					insert_str = arguments[i][9:]
				} else if arguments[i][:2] == "-i" {
					in = true
					insert_str = arguments[i][3:]
				}
			}
		}

		if insert && order || in && o { // sort
			com := order_elem + insert_str

			sort_print(com)
		} else if insert || in {
			com := arguments[1] + insert_str
			for _, char := range com {
				z01.PrintRune(char)
			}
			z01.PrintRune('\n')
		} else if order || o {
			sort_print(order_elem)
		} else {
			for _, char := range arguments[0] {
				z01.PrintRune(char)
			}
			z01.PrintRune('\n')
		}
	}
}

func error() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("	 This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")
}

func sort_print(com string) {
	comb := []rune(com)

	for d := range comb {
		if d < len(comb)-1 {
			for i := 0; i < len(comb)-1; i++ {
				if comb[i] > comb[i+1] {
					comb[i], comb[i+1] = comb[i+1], comb[i]
				}
			}
		}
	}

	for _, char := range comb {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
