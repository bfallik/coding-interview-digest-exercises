package ex09

import (
	"testing"
)

type testInput struct {
	name string
	ns   [][]int
}

func TestSolve(t *testing.T) {
	var tabletests = []struct {
		input  testInput
		output bool
	}{
		{
			input: testInput{
				name: "test 1",
				ns:   [][]int{{1}, {2}, {3}, {}},
			},
			output: true,
		},
		{
			input: testInput{
				name: "test 1",
				ns:   [][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
			},
			output: false,
		},
	}
	for _, tt := range tabletests {
		t.Run(tt.input.name, func(t *testing.T) {
			output := Solve(tt.input.ns)
			if output != tt.output {
				t.Errorf("Expected %v, got %v", tt.output, output)
			}
		})
	}
}
