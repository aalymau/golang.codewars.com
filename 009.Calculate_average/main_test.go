package main

import (
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected float64
	}{
		{
			name:     "Test case #1",
			arr:      []int{1, 2, 3, 4, 5},
			expected: 3.0,
		},
		{
			name:     "Test case #2",
			arr:      []int{},
			expected: 0.0,
		},
		{
			name:     "Test case #3",
			arr:      []int{-1, -2, -3, -4, -5},
			expected: -3.0,
		},
		{
			name:     "Test case #4",
			arr:      []int{-1, 2, 3, 4, -5},
			expected: 0.6,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			println("\nStarting test %s", t.Name())
			// test return code
			result := CalculateAverage(tc.arr)
			if result != tc.expected {
				t.Fatalf("%v got, but %v expected.", result, tc.expected)
			}
		})
	}
}
