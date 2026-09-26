package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		// new data type here
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		// taking the actual input
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) length = %d; want %d", c.input, len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if expectedWord != word {
				t.Errorf("cleanInput(%q)[%d] = %q; want %q", c.input, i, word, expectedWord)
			}
		}
	}
}
