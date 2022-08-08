package piscine

func IterativeFactorial(nb int) int {
	n := 1
	if nb >= 64 {
		n = 0
	} else if nb <= 0 {
		n = 0
	} else {
		for i := nb; i > 0; i-- {
			n = n * i
		}
	}
	return n
}
