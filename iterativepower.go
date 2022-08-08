package piscine

func IterativePower(nb int, power int) int {
	a := 1
	if power < 0 {
		a = 0
	} else {
		for i := 1; i <= power; i++ {
			a *= nb
		}
	}
	return a
}
