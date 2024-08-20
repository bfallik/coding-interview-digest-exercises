package ex04

import "slices"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Solve(heights []int) int {
	volumes := []int{}
	n := len(heights)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			volumes = append(volumes, (j-i)*(min(heights[i], heights[j])))
		}
	}

	return slices.Max(volumes)
}
