package piscine

func StrLen(s string) int {
	count := 0
	for i := range s {
		i++
		count = i
	}
	return count
}
