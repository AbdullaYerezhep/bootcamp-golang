package piscine

func Join(elems []string, sep string) string {
	str:= ""
	for i:=0; i < len(elems); i++{
		str += elems[i]
		if i != len(elems)-1{
		str += sep
		}
	}
	return str
}
