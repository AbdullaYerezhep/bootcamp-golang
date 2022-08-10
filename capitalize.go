package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	str := ""
	j := 0
	for i := 0; i < len(s); i++ {
		j = i
		if j-1 < 0 {
			j = 1
		}
		if runes[i] >= 'a' && runes[i] <= 'z' && ((runes[j-1] >= ' ' && runes[j-1] <= '/') || (runes[j-1] >= ':' && runes[j-1] <= '@')) || i == 0 {
			runes[i] = rune(int(runes[i] - 32))
		}
	}

	for i := 0; i < len(runes); i++ {
		str += string(runes[i])
	}

	return str
}
