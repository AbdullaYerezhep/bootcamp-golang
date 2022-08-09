package piscine

func IsNumeric(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if int(runes[i]) < 48 || int(runes[i]) > 57 {
			return false
		}
	}
	return true
}
