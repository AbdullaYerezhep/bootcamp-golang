package main

import (
	"fmt"
)

func setPoint(ptr []int) []int{
	x:=0
	y:=1
	ptr[x] = 42
	ptr[y] = 21
	return ptr
}

func main() {
	points :=  []int{0, 1}

	points = setPoint(points)

	fmt.Printf("x = %d, y = %d\n",points[0], points[1])
}
