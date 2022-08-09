package piscine

func IsAlpha(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if int(runes[i]) > -1 && int(runes[i]) < 48 || int(runes[i]) > 57 && int(runes[i]) < 65 || int(runes[i]) > 90 && int(runes[i]) < 97 || int(runes[i]) > 123 && int(runes[i]) < 128 {
			return false
		}
	}
	return true
}
