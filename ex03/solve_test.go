package ex03

import (
	"reflect"
	"testing"
)

type testInput struct {
	name  string
	lists [][]int
}

func TestSolve(t *testing.T) {
	var tabletests = []struct {
		input  testInput
		output []int
	}{
		{
			input: testInput{
				name:  "test 1",
				lists: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			},
			output: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			input: testInput{
				name:  "test 2",
				lists: [][]int{},
			},
			output: []int{},
		},
		{
			input: testInput{
				name:  "test 3",
				lists: [][]int{{}},
			},
			output: []int{},
		},
	}
	for _, tt := range tabletests {
		t.Run(tt.input.name, func(t *testing.T) {
			output := Solve(tt.input.lists)
			if !reflect.DeepEqual(output, tt.output) {
				t.Errorf("expected %v, got %v", tt.output, output)
			}
		})
	}
}
