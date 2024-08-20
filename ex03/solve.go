package ex03

func Solve(lists [][]int) []int {
	res := []int{}

	n := 0
	for _, l := range lists {
		n += len(l)
	}

	listsLen := len(lists)
	for i := 0; i < n; i++ {
		min := 0
		minPos := -1
		for j := 0; j < listsLen; j++ {
			if len(lists[j]) == 0 {
				continue
			}
			if minPos == -1 || lists[j][0] < min {
				min = lists[j][0]
				minPos = j
			}
		}

		res = append(res, min)
		lists[minPos] = lists[minPos][1:]
	}

	return res
}
