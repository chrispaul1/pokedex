package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " helLO world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " whats  up    man ",
			expected: []string{"whats", "up", "man"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf(`--------------------------------------------------
			Expecting: %d
			Actual: %d
			Fail, Incorrect length`, len(c.expected), len(actual))
			fmt.Printf("\n")
		} else {
			fmt.Printf(`--------------------------------------------------
			Expecting: %d
			Actual: %d
			Pass, Correct length`, len(c.expected), len(actual))
			fmt.Printf("\n")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf(`--------------------------------------------------
				Expecting: %s
				Actual: %s
				Fail`, word, expectedWord)
				fmt.Printf("\n")
			} else {
				fmt.Printf(`--------------------------------------------------
				Expecting: %s
				Actual: %s
				Pass`, word, expectedWord)
				fmt.Printf("\n")
			}
		}
	}
}
