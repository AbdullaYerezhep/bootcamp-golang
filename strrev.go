package piscine

func StrRev(s string) string {
	chars := []byte(s)
	var revString []rune
	for i := len(chars) - 1; i >= 0; i-- {
		revString = append(revString, rune(chars[i]))
	}
	return string(revString)
}
