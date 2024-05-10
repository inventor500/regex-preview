package regex_preview

import (
	"regexp"
	"testing"
)

func TestRenderOutput(t *testing.T) {
	s := OutputSettings{
		FColor: 90,
		BColor: 52,
	}
	testInputs := []string{
		"hello world",
		"hello world hello",
	}
	expectedResults := []string{
		"\033[38;5;90;48;5;52mhello\033[0m world",
		"\033[38;5;90;48;5;52mhello\033[0m world \033[38;5;90;48;5;52mhello\033[0m",
	}
	// These are the same anyway...
	_regex := regexp.MustCompile("hello")
	testRegex := []*regexp.Regexp{
		_regex,
		_regex,
	}
	for i := 0; i < len(testInputs); i++ {
		output := RenderOutput(testInputs[i], testRegex[i], s)
		if output != expectedResults[i] {
			t.Errorf("Iteration %d: Expected %s, received %s", i, expectedResults[i], output)
		}
	}
}
