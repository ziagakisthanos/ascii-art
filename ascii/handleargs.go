package ascii

import (
	"errors"
	"fmt"
	"strings"
)

func getBannerPath(args []string, banners map[string]string) (string, []string, error) {
	banner := "standard" // Default banner
	if len(args) > 1 {
		lastArg := args[len(args)-1]
		if path, exists := banners[lastArg]; exists {
			return path, args[:len(args)-1], nil
		}
	}
	if path, exists := banners[banner]; exists {
		return path, args, nil
	}
	return "", nil, errors.New(" Invalid banner name. ")
}

func HandleArgs(args []string) error {
	if len(args) < 2 {
		return errors.New(" \nUsage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right \"Hello\" standard\n ")
	}

	// Default alignment and banner
	align := "left"
	banners := map[string]string{
		"standard":   "banners/standard.txt",
		"shadow":     "banners/shadow.txt",
		"thinkertoy": "banners/thinkertoy.txt",
		"other":      "banners/other.txt",
	}

	// Parse alignment
	for i, arg := range args {
		if strings.HasPrefix(arg, "--align=") {
			align = strings.TrimPrefix(arg, "--align=")
			if align != "left" && align != "center" && align != "right" && align != "justify" {
				return fmt.Errorf(" Invalid alignment type: %s\nValid options: left, center, right, justify", align)
			}
			args = append(args[:i], args[i+1:]...)
			break
		} else if arg == "--align" {
			return errors.New(" Missing value for --align flag. Example: --align=right")
		}
	}

	// Parse banner
	bannerPath, updatedArgs, err := getBannerPath(args, banners)
	if err != nil {
		return err
	}
	args = updatedArgs

	// Process input
	input := strings.Join(args[1:], " ")
	input = strings.Replace(input, `\n`, "\n", -1)

	// Generate ASCII art
	asciiWords, _, err := GenerateAsciiArt(input, bannerPath)
	if err != nil {
		return fmt.Errorf(" Error generating ASCII art: %v", err)
	}

	// Align and print
	switch align {
	case "justify":
		for _, line := range justifyWords(asciiWords) {
			fmt.Println(line)
		}
	case "center":
		for _, line := range centerAlign(asciiWords) {
			fmt.Println(line)
		}
	case "right":
		for _, line := range rightAlign(asciiWords) {
			fmt.Println(line)
		}
	default: // Default to left alignment
		for _, line := range leftAlign(asciiWords) {
			fmt.Println(line)
		}
	}

	return nil
}
