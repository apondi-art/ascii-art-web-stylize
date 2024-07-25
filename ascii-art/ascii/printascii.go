// Package ascii provides functions for printing ASCII art
package ascii

import (
	"strings"
)

// PrintArgs contains parameters for the PrintAscii function.
type PrintArgs struct {
	Str        string
	Characters []string
}

// PrintAscii prints ASCII art based on the given PrintArgs configuration.
func PrintAscii(args *PrintArgs) string {
	index := 0
	var result strings.Builder
	// Loop through each line of ASCII art (up to 8 lines)
	for index < 8 {
		for _, char := range args.Str {
			character := args.Characters[int(char)-32]
			lines := strings.Split(character, "\n")
			result.WriteString(lines[index])
		}
		result.WriteString("\n")
		index++
	}
	return result.String()
}
