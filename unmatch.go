package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1; j++ {
			if j == i && j != len(a)-1 {
				j++
			}
			if a[i] == a[j] {
				a[i] = a[len(a)-1]
				a[len(a)-1] = 0
				a = a[:len(a)-1]
			}
			if len(a) == 1 {
				break
			}
		}
	}
	return a[0]
}
