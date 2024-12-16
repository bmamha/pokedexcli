package main
import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	{
		input: "Is this actually, working?",
	  expected: []string{"is", "this", "actually,", "working?"},
	},
}

for _, c := range cases {
	actual := CleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf(`Test Failed:
		input: %v 
	  length of 
		expected: %v  
		and actual: %v 
		does not match`, c.input, len(actual), len(c.expected))
	}
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
	  if word != expectedWord {
			t.Errorf(`
			Test Failed:
			Expected: %v 
			Actual: %v`,expectedWord, word)
		}
	}
}
}
