package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	temp := l

	for i := 0; i < pos; i++ {
		if temp == nil {
			return nil
		} else {
			temp = temp.Next
		}
	}
	return temp
}
