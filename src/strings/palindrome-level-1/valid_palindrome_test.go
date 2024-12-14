package main

import "testing"

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		str      string
		expected bool
	}{
		{"xdd", false},
		{"madam", true},
		{"racecar", true},
		{"hello", false},
		{"", true},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			result := ValidPalindrome(test.str)
			if result != test.expected {
				t.Errorf("For input '%s', expected %v but got %v", test.str, test.expected, result)
			}
		})
	}
}
