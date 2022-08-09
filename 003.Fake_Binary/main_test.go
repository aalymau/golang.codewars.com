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
			coded:    "45385593107843568",
			expected: "01011110001100111",
		},
		{
			name:     "Test case #2",
			coded:    "509321967506747",
			expected: "101000111101101",
		},
		{
			name:     "Test case #3",
			coded:    "366058562030849490134388085",
			expected: "011011110000101010000011011",
		},
		{
			name:     "Test case #4",
			coded:    "15889923",
			expected: "01111100",
		},
		{
			name:     "Test case #5",
			coded:    "800857237867",
			expected: "100111001111",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			println("\nStarting test %s", t.Name())
			// test return code
			result := FakeBin(tc.coded)
			if result != tc.expected {
				t.Fatalf("%s got, but %v expected.", result, tc.expected)
			}
		})
	}
}
