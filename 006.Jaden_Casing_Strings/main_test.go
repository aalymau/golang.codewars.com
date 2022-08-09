package main

import (
	"testing"
)

func TestDecodeMorse(t *testing.T) {
	tests := []struct {
		name     string
		coded    string
		expected string
	}{
		{
			name:     "Test case #1",
			coded:    "most trees are blue",
			expected: "Most Trees Are Blue",
		},
		{
			name:     "Test case #2",
			coded:    "All the rules in this world were made by someone no smarter than you. So make your own.",
			expected: "All The Rules In This World Were Made By Someone No Smarter Than You. So Make Your Own.",
		},
		{
			name:     "Test case #3",
			coded:    "When I die. then you will realize",
			expected: "When I Die. Then You Will Realize",
		},
		{
			name:     "Test case #4",
			coded:    "Jonah Hill is a genius",
			expected: "Jonah Hill Is A Genius",
		},
		{
			name:     "Test case #5",
			coded:    "Dying is mainstream",
			expected: "Dying Is Mainstream",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			println("\nStarting test %s", t.Name())
			// test return code
			result := ToJadenCase(tc.coded)
			if result != tc.expected {
				t.Fatalf("%s got, but %v expected.", result, tc.expected)
			}
		})
	}
}
