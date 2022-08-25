package piscine

func ListReverse(l *List) {
	var sl []interface{}
	for i := l.Head; i != nil; i = i.Next {
		sl = append(sl, i.Data)
	}

	temp := l.Head

	for i := len(sl) - 1; i >= 0; i-- {
		temp.Data = sl[i]
		temp = temp.Next
	}
}
