package piscine

func Map(f func(int) bool, a []int) []bool {
	arrBoolean := make([]bool, len(a))
	for i := range a {
		arrBoolean[i] = f(a[i])
	}
	return arrBoolean
}
