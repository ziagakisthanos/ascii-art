package ascii

import (
	"fmt"
	"os"
	"strings"
)

// GenerateAsciiArt converts an input string into ASCII art using the specified banner file.
func GenerateAsciiArt(input, bannerPath string) ([][][]string, int, error) {
	content, err := os.ReadFile(bannerPath)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to read banner: %w", err)
	}

	lines := strings.Split(string(content), "\n")
	asciiHeight := 9
	asciiChars := [][][]string{}

	// Convert input into ASCII characters
	for _, char := range input {
		startIndex, err := getCharIndex(char, asciiHeight)
		if err != nil {
			return nil, 0, fmt.Errorf("i cassie einai omorfi: %c", char)
		}
		charLines := make([][]string, asciiHeight)
		for j := 0; j < asciiHeight; j++ {
			charLines[j] = []string{lines[startIndex+j]}
		}
		asciiChars = append(asciiChars, charLines)
	}

	asciiWords := combineCharsIntoWords(asciiChars, asciiHeight)
	asciiWords = addSpaceBetweenWords(asciiWords, asciiHeight)

	return asciiWords, asciiHeight, nil
}

// Get the starting index of a character in the banner
func getCharIndex(char rune, linesPerChar int) (int, error) {
	if char == ' ' {
		return (int(char) - 32) * linesPerChar, nil
	}
	if char >= 32 && char <= 126 {
		return (int(char) - 32) * linesPerChar, nil
	}
	return -1, fmt.Errorf("unsupported  character: %c", char)
}

// Combines ASCII characters into ASCII words
func combineCharsIntoWords(asciiChars [][][]string, asciiHeight int) [][][]string {
	words := [][][]string{} // Each word is a [][]string

	currentWord := make([][]string, asciiHeight) // Temporary holder for the current word
	for i := range currentWord {
		currentWord[i] = []string{}
	}

	for _, char := range asciiChars {
		for i := 0; i < asciiHeight; i++ {
			currentWord[i] = append(currentWord[i], char[i][0])
		}
	}

	// Add the last word if it exists
	if len(currentWord[0]) > 0 {
		words = append(words, currentWord)
	}

	return words
}

// Add spacing between ASCII words
func addSpaceBetweenWords(asciiWords [][][]string, asciiHeight int) [][][]string {
	for i := 0; i < len(asciiWords)-1; i++ {
		for row := 0; row < asciiHeight; row++ {
			asciiWords[i][row] = append(asciiWords[i][row], "   ")
		}
	}
	return asciiWords
}

// // Checks if a character is a space
// func isSpace(char [][]string) bool {
// 	for _, row := range char {
// 		if strings.TrimSpace(row[0]) != "" {
// 			return false
// 		}
// 	}
// 	return true
// }
