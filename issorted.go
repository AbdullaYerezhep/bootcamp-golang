package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := 0
	descending := 0
	equal := 0
	for i := range a {
		if i != len(a)-1 {
			if f(a[i], a[i+1]) < 0 {
				ascending++
			} else if f(a[i], a[i+1]) > 0 {
				descending++
			} else {
				equal++
			}
		}
	}

	if ascending == len(a)-1 || descending == len(a)-1 || equal == len(a)-1 {
		return true
	}
	return false
}
