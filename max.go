package piscine

func Max(a []int) int {
	max := -99
	for i := 0; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}
