package piscine

func AlphaCount(s string) int {
	runes := []rune(s)
	count := 0
	for i:=0; i<len(s); i++ {
		if(int(runes[i]) > 64 && int(runes[i]) < 91 ||  int(runes[i]) > 96 && int(runes[i]) < 123) {
			count++
		}
	}
	return count
}
