package regex_preview

import (
	"fmt"
	"regexp"
	"strings"
)

type OutputSettings struct {
	FColor int // Foreground
	BColor int // Background
}

func (s OutputSettings) String() string {
	return fmt.Sprintf("\033[38;5;%d;48;5;%dm", s.FColor, s.BColor)
}

func RenderOutput(content string, regex *regexp.Regexp, settings OutputSettings) string {
	// Generate matches
	matches := regex.FindAllStringIndex(content, -1)
	if matches == nil { // Nothing matches
		return content
	}
	builder := strings.Builder{}
	current := 0 // Track the end of the previous match
	for match := 0; match < len(matches); match++ {
		start := matches[match][0]
		end := matches[match][1]
		// Write until the start
		builder.WriteString(content[current:start])
		// Write the colored output
		builder.WriteString(settings.String())
		builder.WriteString(content[start:end])
		// End the colored output
		builder.WriteString("\033[0m")
		current = end
	}
	// Write whatever remains
	builder.WriteString(content[current:])
	return builder.String()
}
