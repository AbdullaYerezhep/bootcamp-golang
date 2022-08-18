package piscine

func Abort(a, b, c, d, e int) int {
	var median int
	arr := []int{a, b, c, d, e}
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j+1] > arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	length := len(arr)
	median = arr[length/2]
	return median
}
