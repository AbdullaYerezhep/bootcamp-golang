package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	str := ""
	isnew := true
	for i := 0; i < len(s); i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' || runes[i] >= '0' && runes[i] <= '9' || runes[i] >= 'A' && runes[i] <= 'Z' {
			if runes[i] >= 'A' && runes[i] <= 'Z' && !isnew {
				runes[i] = rune(int(runes[i] + 32))
			} else if runes[i] >= 'a' && runes[i] <= 'z' && isnew {
				runes[i] = rune(int(runes[i] - 32))
			}
			isnew = false
		} else {
			isnew = true
		}
	}

	for i := 0; i < len(runes); i++ {
		str += string(runes[i])
	}

	return str
}
