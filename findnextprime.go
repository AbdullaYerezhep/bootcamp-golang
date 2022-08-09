package piscine

func FindNextPrime(nb int) int {
	if nb <= 0 || nb == 1 {
		return 2
	}

	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return FindNextPrime(nb + 1)
		}
	}

	return nb
}
