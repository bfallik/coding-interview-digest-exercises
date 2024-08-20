package ex06

import (
	"testing"
)

type testInput struct {
	name string
	list []int
}

func TestSolve(t *testing.T) {
	var tabletests = []struct {
		input testInput
	}{
		{
			input: testInput{
				name: "test 1",
				list: []int{1, 2, 3, 1},
			},
		},
		{
			input: testInput{
				name: "test 2",
				list: []int{1, 2, 1, 3, 5, 6, 4},
			},
		},
	}
	for _, tt := range tabletests {
		t.Run(tt.input.name, func(t *testing.T) {
			output := Solve(tt.input.list)
			if !isValidPeak(tt.input.list, output) {
				t.Errorf("invalid peak %v", output)
			}
		})
	}
}
