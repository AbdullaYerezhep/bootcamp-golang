package piscine

func IterativeFactorial(nb int) int {
	n := 1
	if nb < 64 && nb > 0 {
		for i := nb; i > 0; i-- {
			n = n * i
		}
		return n
	}else{
		return 0
	}
}
