package piscine

func ListClear(l *List) {
	if l.Head != nil {
		l.Head = nil
	}
	if l.Tail != nil {
		l.Tail = nil
	}
}
