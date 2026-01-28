package main
import(
	"testing"
)

func TestCleanInput(t *testing.T){
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  small  ",
			expected: []string{"small"},
		},
		{
			input:    "  small words  ",
			expected: []string{"small", "words"},
		},
		{
			input:    "  BIG WORDS  ",
			expected: []string{"big", "words"},
		},
	}

		for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}