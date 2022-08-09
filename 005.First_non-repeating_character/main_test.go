package main

import (
	"testing"
)

func TestMaps(t *testing.T) {
	tests := []struct {
		name     string
		initial  string
		expected string
	}{
		{
			name:     "Test case #1",
			initial:  "a",
			expected: "a",
		},
		{
			name:     "Test case #2",
			initial:  "stress",
			expected: "t",
		},
		{
			name:     "Test case #3",
			initial:  "moonmen",
			expected: "e",
		},
		{
			name:     "Test case #4",
			initial:  "",
			expected: "",
		},
		{
			name:     "Test case #5",
			initial:  "abba",
			expected: "",
		},
		{
			name:     "Test case #6",
			initial:  "aa",
			expected: "",
		},
		{
			name:     "Test case #7",
			initial:  "~><#~><",
			expected: "#",
		},
		{
			name:     "Test case #8",
			initial:  "hello world, eh?",
			expected: "w",
		},
		{
			name:     "Test case #9",
			initial:  "sTreSS",
			expected: "T",
		},
		{
			name:     "Test case #10",
			initial:  "Go hang a salami, I'm a lasagna hog!",
			expected: ",",
		}}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			println("\nStarting test %s", t.Name())
			result := FirstNonRepeating(tc.initial)

			// Check result
			if result != tc.expected {
				t.Fatalf("%v got, but %v expected.", result, tc.initial)
			}
		})
	}
}
