package piscine

func StrRev(s string) {
	chars := []rune(s)
	revString := ""
	for i := len(chars); i > 0; i-- {
		revString += string(chars[i])
	}
	s = revString
}
