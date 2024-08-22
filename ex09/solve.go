package ex09

func SolveRoom(rooms [][]int, visited []int, room int) {
	visited[room] = 1

	keys := rooms[room]
	for _, key := range keys {
		if visited[key] == 0 {
			SolveRoom(rooms, visited, key)
		}
	}
}

func Solve(rooms [][]int) bool {
	visited := make([]int, len(rooms))
	SolveRoom(rooms, visited, 0)

	for _, v := range visited {
		if v == 0 {
			return false
		}
	}
	return true
}
