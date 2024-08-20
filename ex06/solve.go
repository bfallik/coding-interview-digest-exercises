package ex06

func isValidPeak(list []int, pos int) bool {
	n := len(list)
	if pos == 0 {
		return list[0] > list[1]
	} else if pos == n-1 {
		return list[n-1] > list[n-2]
	} else {
		return list[pos] > list[pos-1] && list[pos] > list[pos+1]
	}
}

func Solve(list []int) int {
	for i := range list {
		if isValidPeak(list, i) {
			return i
		}
	}
	panic("not reached")
}
