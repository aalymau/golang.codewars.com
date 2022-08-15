package main

import (
	"testing"
)

func TestPositiveSum(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "Test case #1",
			arr:      []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "Test case #2",
			arr:      []int{1, -2, 3, 4, 5},
			expected: 13,
		},
		{
			name:     "Test case #3",
			arr:      []int{},
			expected: 0,
		},
		{
			name:     "Test case #4",
			arr:      []int{-1, -2, -3, -4, -5},
			expected: 0,
		},
		{
			name:     "Test case #5",
			arr:      []int{-1, 2, 3, 4, -5},
			expected: 9,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			println("\nStarting test %s", t.Name())
			// test return code
			result := PositiveSum(tc.arr)
			if result != tc.expected {
				t.Fatalf("%v got, but %v expected.", result, tc.expected)
			}
		})
	}
}
