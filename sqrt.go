package piscine

func Sqrt(nb int) int {
	m := 0
	for i := 1; i < nb; i++ {
		if i*i == nb {
			m = i
		}
	}
	if m % 1 != 0{
		return 0
	}else{
		return m
	}
	
}
