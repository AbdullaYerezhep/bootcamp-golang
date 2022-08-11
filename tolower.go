package piscine

func ToLower(s string) string {
	runes := []rune(s)
	var upperrunes []rune
	for i := 0; i < len(s); i++ {
		if int(runes[i]) > 64 && int(runes[i]) < 91 {
			upperrunes = append(upperrunes, rune(int(runes[i])+32))
		} else {
			upperrunes = append(upperrunes, rune(int(runes[i])))
		}
	}
	return string(upperrunes)
}
