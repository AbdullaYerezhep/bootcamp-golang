package piscine

func ActiveBits(n int) int {
	array := []int{}

	for n != 0 {

		array = append(array, n%2)
		n = n / 2
	}
	add := 0
	for i := 0; i < len(array); i++ {
		if array[i] == 1 {
			add++
		}
	}
	return add
}
