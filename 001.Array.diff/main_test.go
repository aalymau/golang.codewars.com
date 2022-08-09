package main

import (
	"testing"
)

func TestArrayDiff(t *testing.T) {
	tests := []struct {
		name     string
		arr_a    []int
		arr_b    []int
		expected []int
	}{
		{
			name:     "Test case #1",
			arr_a:    []int{1, 2},
			arr_b:    []int{1},
			expected: []int{2},
		},
		{
			name:     "Test case #2",
			arr_a:    []int{1, 2, 2},
			arr_b:    []int{1},
			expected: []int{2, 2},
		},
		{
			name:     "Test case #3",
			arr_a:    []int{1, 2, 2},
			arr_b:    []int{2},
			expected: []int{1},
		},
		{
			name:     "Test case #4",
			arr_a:    []int{1, 2, 2},
			arr_b:    []int{},
			expected: []int{1, 2, 2},
		},
		{
			name:     "Test case #5",
			arr_a:    []int{},
			arr_b:    []int{1, 2},
			expected: []int{},
		},
		{
			name:     "Test case #6",
			arr_a:    []int{1, 2, 3},
			arr_b:    []int{1, 2},
			expected: []int{3},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			println("\nStarting test %s", t.Name())
			// test return code
			result := ArrayDiff(tc.arr_a, tc.arr_b)
			if len(result) != len(tc.expected) {
				t.Fatalf("Array lenght: %v got, but %v expected", len(result), len(tc.expected))
			}
			for i, val := range tc.expected {
				if result[i] != val {
					t.Fatalf("Array[%v]: %v got, but %v expected.", i, result[i], val)
				}
			}
		})
	}
}
