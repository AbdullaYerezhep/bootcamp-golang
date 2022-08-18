package piscine

func Rot14(s string) string {
	runes := []rune(s)
	result := ""
	for i, char := range runes {
		if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' {
			if char <= 'L' && char >= 'A' || char <= 'l' && char >= 'a' {
				runes[i] = rune(int(char) + 14)
			} else {
				runes[i] = rune(int(char) - 12)
			}
		}
	}

	for i := range runes {
		result += string(runes[i])
	}
	return result
}
