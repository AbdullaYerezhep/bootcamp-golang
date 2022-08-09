package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	var upperrunes []rune
	for i := 0; i < len(s); i++ {
		if int(runes[i]) > 96 && int(runes[i]) < 123{
			upperrunes = append(upperrunes, rune(int(runes[i])-32))
		} else {
			upperrunes = append(upperrunes, rune(int(runes[i])))
		}
	}
	return string(upperrunes)
}
