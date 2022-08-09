package piscine

func FindNextPrime(nb int) int {
	result := 1
	if nb <= 0 || nb == 1 {
		result = 2
	} else {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				result = 0
			}
		}

		if result == 0 {
			nb++
			for j := 2; j < nb; j++ {
				if nb%j == 0 {
					nb++
					j = 2
				}
			}
			result = nb
		} else {
			result = nb
		}
	}
	return result
}
