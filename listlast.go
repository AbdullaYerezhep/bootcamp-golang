package piscine

func ListLast(l *List) interface{} {
	if l.Tail != nil {
		return l.Tail.Data
	} else {
		return nil
	}
}
