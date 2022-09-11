package model

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type DivinerSpell struct {
	gorm.Model

	DivinationSchool string
	Tier             string
	Name             string
	Description      string
	Range            string
	CastingTime      string
	Duration         string
	AreaOfEffect     string
	Damage           string
	SavingThrow      string
	Reversible       string
}

const maxDivinerLineLength = 85

func (s *DivinerSpell) PrintSpell() string {

	sb := fmt.Sprintf(
		"\n\n\n\n\n"+
			"Divination school: %s\n"+
			"Tier: %s\n"+
			"Name: %s\n"+
			"Range: %s\n"+
			"CastingTime: %s\n"+
			"Duration: %s\n"+
			"AreaOfEffect: %s\n"+
			"Damage: %s\n"+
			"SavingThrow: %s\n"+
			"Reversible: %s\n"+
			"Description:\n\n%s\n",
		s.DivinationSchool,
		s.Tier,
		s.Name,
		s.Range,
		s.CastingTime,
		s.Duration,
		s.AreaOfEffect,
		s.Damage,
		s.SavingThrow,
		s.Reversible,
		s.formatDescription(),
	)

	return sb
}

func (s *DivinerSpell) formatDescription() string {
	currentLineLength := 0

	var result string
	splitString := strings.Split(s.Description, " ")

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

		if currentLineLength+lenFirstWord <= maxDivinerLineLength {
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
