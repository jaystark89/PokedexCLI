package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "hello world",
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "Go IS FUN",
			input:    "Go IS FUN",
			expected: []string{"go", "is", "fun"},
		},
	}
	// add more cases here

	// loop over the cases

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			// error and continue here

			t.Errorf("name: %s, expected: %v, got: %v", c.name, c.expected, actual)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			//Check each word in the slice
			//if they don't match, use t.Errorf to print an error message
			// and fail the test

			if word != expectedWord {
				t.Errorf("name: %s, expected: %v, got: %v", c.name, expectedWord, word)
				continue
			}

		}
	}

}
