package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	count := 0
	for i := range a {
		if i != len(a)-1 {
			if f(a[i], a[i+1]) < 0 {
				count++
			}
		}
	}

	return count == len(a)-1
}
