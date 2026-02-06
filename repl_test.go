
package main

import "testing"

func TestCleanInput(t *testing.T) {
    cases := []struct{
		input string
		expected []string
	}{
		{
			input: " Hello,    World!  ",
			expected: []string{"hello,", "world!"},
		},
		{
			input: "goTtA caTch em' aLl!",
			expected: []string{"gotta", "catch", "em'", "all!"},
		},
		{
			input: "",
			expected: []string{},
		},
				{
			input: "      ",
			expected: []string{},
		},
						{
			input: "	",
			expected: []string{},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
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