package piscine

func Index(s string, toFind string) int {
	if len(s) == 0 {
		return 0
	} else if len(toFind) == 0 {
		return 0
	}
	q := 0
	for i := range s {
		if s[i] == toFind[q] {
			if q == len(toFind)-1 {
				return i - (len(toFind) - 1)
			}
			q++
		}
	}
	return -1
}
