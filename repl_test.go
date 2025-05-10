package main

import "testing"

func TestCleanInput(t *testing.T) {	
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
	}

	

	for _,c := range cases{
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("Length of output does not match expected slice length of %v",c.expected)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord{
				t.Errorf("%v does not match expected word of %v",word,expectedWord)
			}
		}
	}
}