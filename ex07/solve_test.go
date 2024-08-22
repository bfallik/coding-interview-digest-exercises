package ex07

import (
	"slices"
	"testing"
)

type testInput struct {
	name string
	ns   []int
	k    int
}

func TestSolve(t *testing.T) {
	var tabletests = []struct {
		input  testInput
		output []int
	}{
		{
			input: testInput{
				name: "test 1",
				ns:   []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			output: []int{1, 2},
		},
		{
			input: testInput{
				name: "test 2",
				ns:   []int{1},
				k:    1,
			},
			output: []int{1},
		},
	}
	for _, tt := range tabletests {
		t.Run(tt.input.name, func(t *testing.T) {
			output := Solve(tt.input.ns, tt.input.k)
			slices.Sort(output)
			if !slices.Equal(output, tt.output) {
				t.Errorf("Expected %v, got %v", tt.output, output)
			}
		})
	}
}
