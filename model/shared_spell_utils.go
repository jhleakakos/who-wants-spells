package model

import "strings"

const maxLineLength = 85

func FormatDescription(description string) string {
	currentLineLength := 0

	var result string
	splitString := strings.Split(description, " ")

	for _, word := range splitString {

		newLineSplit := []string{"", ""}

		if strings.Contains(word, "\n\n") {
			newLineIndex := strings.Index(word, "\n")
			newLineSplit[0] = word[:newLineIndex]
			newLineSplit[1] = word[newLineIndex+1:]
		} else if strings.Contains(word, "\n") {
			newLineSplit = strings.Split(word, "\n")
		} else {
			newLineSplit[0] = word
		}

		lenFirstWord := len(newLineSplit[0])
		lenSecondWord := len(newLineSplit[1])

		if currentLineLength+lenFirstWord <= maxLineLength {
			result += newLineSplit[0]
			result += " "
			currentLineLength += lenFirstWord + 1
		} else {
			result += "\n"
			result += newLineSplit[0]
			result += " "
			currentLineLength = lenFirstWord + 1
		}

		if newLineSplit[1] != "" {
			result += "\n"
			result += newLineSplit[1]
			result += " "
			currentLineLength = lenSecondWord + 1
		}
	}

	return result
}
