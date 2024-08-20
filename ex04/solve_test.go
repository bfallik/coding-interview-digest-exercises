package ex04

import (
	"testing"
)

type testInput struct {
	name    string
	heights []int
}

func TestSolve(t *testing.T) {
	var tabletests = []struct {
		input  testInput
		output int
	}{
		{
			input: testInput{
				name:    "test 1",
				heights: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			output: 49,
		},
		{
			input: testInput{
				name:    "test 2",
				heights: []int{1, 1},
			},
			output: 1,
		},
	}
	for _, tt := range tabletests {
		t.Run(tt.input.name, func(t *testing.T) {
			output := Solve(tt.input.heights)
			if output != tt.output {
				t.Errorf("expected %v, got %v", tt.output, output)
			}
		})
	}
}
