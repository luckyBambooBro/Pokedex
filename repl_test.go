package main

import "testing"

func TestCleanInput(t *testing.T) {
    cases := []struct{
		input string
		expected []string
	}{
		{
			input: "Hello World!",
			expected: []string{"Hello", "World!"},
		},
		{
			input: "gotta catch em' all!",
			expected: []string{"gotta", "catch", "em'", "all"},
		},
		{
			input: "",
			expected: []string{""},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length mismatch: got %d, want %d", len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word mismatch at index %d. Got %q, want %q", i, actual[i], c.expected[i])
			}
		}
	}
}