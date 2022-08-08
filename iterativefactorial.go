package piscine

func IterativeFactorial(nb int) int {
	n := 1
	if nb >= 64 && nb <= 0 {
		return 0
	} else {
		for i := nb; i > 0; i-- {
			n = n * i
		}
		return n
	}
}
