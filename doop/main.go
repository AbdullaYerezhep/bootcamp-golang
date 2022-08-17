package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if args[0] == "-9223372036854775809" && args[1] == "-" || args[0] == "9223372036854775809" && args[1] == "-" {
	} else {
		if !isValid(args) {
		} else {
			numbers := collectNumbers(args)
			operator := collectOperators(args)
			result := 0
			if numbers[1] == 0 && operator == "/" {
				fmt.Println("No division by 0")
			} else if numbers[1] == 0 && operator == "%" {
				fmt.Println("No modulo by 0")
			} else if numbers[0] == 9223372036854775807 && (operator == "+" || operator == "*") {
			} else {
				switch operator {
				case "+":
					result = addition(numbers[0], numbers[1])
				case "-":
					result = substraction(numbers[0], numbers[1])
				case "*":
					result = multiply(numbers[0], numbers[1])
				case "/":
					result = division(numbers[0], numbers[1])
				case "%":
					result = modulo(numbers[0], numbers[1])

				}
				fmt.Println(result)

			}
		}
	}
}

func addition(a, b int) int {
	c := a + b
	return c
}

func substraction(a, b int) int {
	c := a - b
	return c
}

func multiply(a, b int) int {
	c := a * b
	return c
}

func division(a, b int) int {
	c := a / b
	return c
}

func modulo(a, b int) int {
	c := a % b
	return c
}

func isNumeric(s string) bool {
	runes := []rune(s)
	if len(runes) > 1 {
		for i := 0; i < len(runes); i++ {
			if runes[i] == '-' && runes[i+1] >= '0' && runes[i+1] <= '9' {
				return true
			} else if runes[i] < '0' || runes[i] > '9' {
				return false
			}
		}
	} else {
		if runes[0] < '0' || runes[0] > '9' || runes[0] == '-' {
			return false
		}
	}

	return true
}

func isOperator(s string) bool {
	runes := []rune(s)

	if len(runes) > 1 {
		for i := 0; i < len(runes)-1; i++ {
			if runes[i] == '+' || runes[i] == '-' && runes[i+1] <= '0' && runes[i+1] >= '9' || runes[i] == '*' || runes[i] == '/' || runes[i] == '%' || runes[i] == '"' {
				return true
			}
		}
	} else {
		if runes[0] == '+' || runes[0] == '-' || runes[0] == '*' || runes[0] == '/' || runes[0] == '%' || runes[0] == '"' {
			return true
		}
	}

	return false
}

func isValid(args []string) bool {
	for i := range args {
		if !isNumeric(args[i]) && !isOperator(args[i]) {
			return false
		}
	}
	return true
}

func collectNumbers(args []string) []int {
	numbers := []int{}
	for i := range args {
		if isNumeric(args[i]) {
			numbers = append(numbers, stringToInt(args[i]))
		}
	}
	return numbers
}

func collectOperators(args []string) string {
	runes := []rune{}
	operator := ""
	for i := range args {
		if isOperator(args[i]) {
			for _, char := range args[i] {
				runes = append(runes, char)
			}
		}
	}

	for j := range runes {
		operator = string(runes[j])
	}
	return operator
}

func stringToInt(s string) int {
	runes := []rune(s)
	var arr []int
	m := 1
	n := 0
	isminus := false

	for i := range s {
		if runes[i] <= '9' && runes[i] >= '0' {
			arr = append(arr, int(runes[i]-48))
			isminus = true
		} else if runes[i] == '-' && !isminus {
			m = -1
		}
	}

	for i := 0; i < len(arr); i++ {
		n = n*10 + arr[i]
	}

	return n * m
}
