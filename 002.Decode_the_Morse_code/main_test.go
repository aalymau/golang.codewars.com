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
			coded:    ".... . -.--   .--- ..- -.. .",
			expected: "HEY JUDE",
		},
		{
			name:     "Test case #2",
			coded:    ".-",
			expected: "A",
		},
		{
			name:     "Test case #3",
			coded:    ".",
			expected: "E",
		},
		{
			name:     "Test case #4",
			coded:    "..",
			expected: "I",
		},
		{
			name:     "Test case #5",
			coded:    ". .",
			expected: "EE",
		},
		{
			name:     "Test case #6",
			coded:    ".   .",
			expected: "E E",
		},
		{
			name:     "Test case #7",
			coded:    "...---...",
			expected: "SOS",
		},
		{
			name:     "Test case #8",
			coded:    "... --- ...",
			expected: "SOS",
		},
		{
			name:     "Test case #9",
			coded:    "...   ---   ...",
			expected: "S O S",
		},
		{
			name:     "Test case #10",
			coded:    "      ...---... -.-.--   - .... .   --.- ..- .. -.-. -.-   -... .-. --- .-- -.   ..-. --- -..-   .--- ..- -- .--. ...   --- ...- . .-.   - .... .   .-.. .- --.. -.--   -.. --- --. .-.-.-  ",
			expected: "SOS! THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG.",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			println("\nStarting test %s", t.Name())
			// test return code
			result := DecodeMorse(tc.coded)
			if result != tc.expected {
				t.Fatalf("%s got, but %v expected.", result, tc.expected)
			}
		})
	}
}
