package ex05

import (
	"testing"
)

type testInput struct {
	name string
	head []int
	pos  int
}

func TestSolve(t *testing.T) {
	var tabletests = []struct {
		input  testInput
		output string
	}{
		{
			input: testInput{
				name: "test 1",
				head: []int{3, 2, 0, -4},
				pos:  1,
			},
			output: "index 1",
		},
		{
			input: testInput{
				name: "test 2",
				head: []int{1, 2},
				pos:  0,
			},
			output: "index 0",
		},
		{
			input: testInput{
				name: "test 3",
				head: []int{1},
				pos:  -1,
			},
			output: "null",
		},
	}
	for _, tt := range tabletests {
		t.Run(tt.input.name, func(t *testing.T) {
			var l List
			for _, n := range tt.input.head {
				l.PushBack(n)
			}
			if tt.input.pos != -1 {
				l.AddBacklink(tt.input.pos)
			}
			output := Solve(l)
			if output != tt.output {
				t.Errorf("expected %v, got %v", tt.output, output)
			}
		})
	}
}
