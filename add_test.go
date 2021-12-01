package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(2, 2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input1   int
		input2   int
		expected int
	}{
		{2, 2, 4},
		{-1, 1, 0},
		{0, 2, 22},
	}

	for _, test := range tests {
		if output := Add(test.input1, test.input2); output != test.expected {
			t.Error("Test Failed:", test.input1, test.input2, test.expected, output)
		}
	}
}

// for coverage:
// go test -cover
