package main

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name     string
		coded    string
		expected string
	}{
		{
			name:     "Test case #1",
			coded:    "The quick brown fox jumps over the lazy dog.",
			expected: "ehT kciuq nworb xof spmuj revo eht yzal .god",
		},
		{
			name:     "Test case #2",
			coded:    "apple",
			expected: "elppa",
		},
		{
			name:     "Test case #3",
			coded:    "a b c d",
			expected: "a b c d",
		},
		{
			name:     "Test case #4",
			coded:    "double  spaced  words",
			expected: "elbuod  decaps  sdrow",
		},
		{
			name:     "Test case #5",
			coded:    "stressed desserts",
			expected: "desserts stressed",
		},
		{
			name:     "Test case #6",
			coded:    "hello hello",
			expected: "olleh olleh",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			println("\nStarting test %s", t.Name())
			// test return code
			result := ReverseWords(tc.coded)
			if result != tc.expected {
				t.Fatalf("%s got, but %v expected.", result, tc.expected)
			}
		})
	}
}
