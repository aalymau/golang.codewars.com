package main

import (
	"testing"
)

func TestMaps(t *testing.T) {
	tests := []struct {
		name     string
		arr_a    []int
		expected []int
	}{
		{
			name:     "Test case #1",
			arr_a:    []int{1, 2, 3},
			expected: []int{2, 4, 6},
		},
		{
			name:     "Test case #2",
			arr_a:    []int{4, 1, 1, 1, 4},
			expected: []int{8, 2, 2, 2, 8},
		},
		{
			name:     "Test case #3",
			arr_a:    []int{2, 2, 2, 2, 2, 2},
			expected: []int{4, 4, 4, 4, 4, 4},
		},
		{
			name:     "Test case #4",
			arr_a:    []int{},
			expected: []int{},
		},
		{
			name:     "Test case #5",
			arr_a:    []int{},
			expected: []int{},
		},
		{
			name:     "Test case #6",
			arr_a:    []int{},
			expected: []int{},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			println("\nStarting test %s", t.Name())
			result := Maps(tc.arr_a)

			// Check content of initial map
			for i, val := range tc.arr_a {
				if tc.arr_a[i] != val {
					t.Fatalf("Initial array[%v]: %v got, but %v expected.", i, tc.arr_a[i], val)
				}
			}
			// Check size of map
			if len(result) != len(tc.expected) {
				t.Fatalf("Array lenght: %v got, but %v expected", len(result), len(tc.expected))
			}
			// Check content of result map
			for i, val := range tc.expected {
				if result[i] != val {
					t.Fatalf("Array[%v]: %v got, but %v expected.", i, result[i], val)
				}
			}
		})
	}
}
