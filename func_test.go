package main
import(
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
			input:    "  Charmander Bulbasaur PIKACHU   ",
			expected: []string{"charmander", "bulbasaur","pikachu"},
		},
		{
			input:    "  Todadile Meowth Mewtwo  ",
			expected: []string{"todadile", "meowth","mewtwo"},
		},
	}	

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("Results are incorrect length: a:%d e:%d",len(actual),len(c.expected))
			return
		}
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Results are do not match: a:%v e:%v",word,expectedWord)
				return
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
	return
}