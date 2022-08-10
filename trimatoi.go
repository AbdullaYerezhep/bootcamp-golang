package piscine

func TrimAtoi(s string) int {
	runes := []rune(s)
	var arr []int
	n := 0
	m := 1
	isminus := false
	for i := range s {
		if runes[i] <= '9' && runes[i] >= '0' {
			arr = append(arr, int(runes[i]-48))
			isminus = true
		} else if runes[i] == '-' && !isminus {
			m = -1
		}
	}

	for i := 0; i < len(arr); i++ {
		n = n*10 + arr[i]
	}

	return n * m
}
