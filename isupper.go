package piscine

func IsUpper(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if int(runes[i]) < 97 || int(runes[i]) > 91 {
			return false
		}
	}
	return true
}
