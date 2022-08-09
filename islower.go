package piscine

func IsLower(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if int(runes[i]) < 97 || int(runes[i]) > 122 {
			return false
		}
	}
	return true
}
