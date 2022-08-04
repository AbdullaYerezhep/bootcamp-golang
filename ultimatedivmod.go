package piscine

func UltimateDivMod(a *int, b *int) {
	d := *a / *b
	r := *a % *b
	*a = d
	*b = r
}
