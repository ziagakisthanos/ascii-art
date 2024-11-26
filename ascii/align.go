package ascii

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GetTerminalWidth dynamically fetches the width of the terminal
func GetTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 80 // Default width if terminal size retrieval fails
	}
	var rows, cols int
	fmt.Sscanf(string(out), "%d %d", &rows, &cols)
	return cols
}

// justifyWords justifies ASCII words to fill the terminal width.
func justifyWords(asciiWords [][][]string) []string {
	if len(asciiWords) == 0 {
		return []string{}
	}

	terminalWidth := GetTerminalWidth()
	numRows := len(asciiWords[0])
	var result []string

	for row := 0; row < numRows; row++ {
		// Calculate total width and spaces
		totalWordLength := 0
		for _, word := range asciiWords {
			totalWordLength += len(strings.Join(word[row], ""))
		}

		totalSpaces := terminalWidth - totalWordLength
		if totalSpaces < 0 {
			totalSpaces = 0
		}

		numGaps := len(asciiWords) - 1
		spacePerGap := 0
		extraSpaces := 0
		if numGaps > 0 {
			spacePerGap = totalSpaces / numGaps
			extraSpaces = totalSpaces % numGaps
		}

		// Build justified line
		var line string
		for i, word := range asciiWords {
			line += strings.Join(word[row], "")
			if i < numGaps {
				spaces := spacePerGap
				if extraSpaces > 0 {
					spaces++
					extraSpaces--
				}
				line += strings.Repeat(" ", spaces)
			}
		}

		result = append(result, line)
	}
	return result
}

// centerAlign centers ASCII words in the terminal.
func centerAlign(asciiWords [][][]string) []string {
	terminalWidth := GetTerminalWidth()
	numRows := len(asciiWords[0])
	var result []string

	for row := 0; row < numRows; row++ {
		var line string
		for _, word := range asciiWords {
			line += strings.Join(word[row], "")
		}

		// Calculate padding
		padding := (terminalWidth - len(line)) / 2
		if padding < 0 {
			padding = 0
		}

		result = append(result, strings.Repeat(" ", padding)+line)
	}

	return result
}

// leftAlign aligns the ASCII art to the left
func leftAlign(asciiWords [][][]string) []string {
	numRows := len(asciiWords[0])
	var result []string

	for row := 0; row < numRows; row++ {
		var line string
		for _, word := range asciiWords {
			line += strings.Join(word[row], "")
		}
		result = append(result, line)
	}

	return result
}

// rightAlign aligns the ASCII art to the right
func rightAlign(asciiWords [][][]string) []string {
	terminalWidth := GetTerminalWidth()
	numRows := len(asciiWords[0])
	var result []string

	for row := 0; row < numRows; row++ {
		var line string
		for _, word := range asciiWords {
			line += strings.Join(word[row], "")
		}
		padding := terminalWidth - len(line)
		result = append(result, strings.Repeat(" ", padding)+line)
	}

	return result
}
