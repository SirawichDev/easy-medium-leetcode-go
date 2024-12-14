package main

import "testing"

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		str      string
		expected bool
	}{
		{"{}", true}, {"{", false},
		{"[{}]", true}, {"{{}", false},
		{"[[", false}, {"xx", false},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			result := ValidParentheses(test.str)
			if result != test.expected {
				t.Errorf("For input '%s', expected %v but got %v", test.str, test.expected, result)
			}
		})
	}
}
