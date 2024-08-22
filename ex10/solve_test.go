package ex10

import (
	"testing"
)

type testInput struct {
	name string
	m    int
	n    int
}

func TestSolve(t *testing.T) {
	var tabletests = []struct {
		input  testInput
		output int
	}{
		{
			input: testInput{
				name: "test 1",
				m:    3,
				n:    7,
			},
			output: 28,
		},
		{
			input: testInput{
				name: "test 1",
				m:    3,
				n:    2,
			},
			output: 3,
		},
	}
	for _, tt := range tabletests {
		t.Run(tt.input.name, func(t *testing.T) {
			output := Solve(tt.input.m, tt.input.n)
			if output != tt.output {
				t.Errorf("Expected %v, got %v", tt.output, output)
			}
		})
	}
}
