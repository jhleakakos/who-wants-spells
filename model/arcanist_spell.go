package model

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type ArcanistSpell struct {
	gorm.Model
	Level             string
	SpellName         string
	Frequency         string
	Description       string
	Range             string
	CastingTime       string
	Duration          string
	AreaOfEffect      string
	Damage            string
	SavingThrow       string
	Reversible        string
	Components        string
	SpecialComponents string
	Class             string
}

const maxLineLength = 85

func (s *ArcanistSpell) PrintSpell() string {

	sb := fmt.Sprintf(
		"\n\n\n\n\n"+
			"Spell Name: %s\n"+
			"Level: %s\n"+
			"Frequency: %s\n"+
			"Range: %s\n"+
			"Casting Time: %s\n"+
			"Duration: %s\n"+
			"Area of Effect: %s\n"+
			"Damage: %s\n"+
			"Saving Throw: %s\n"+
			"Reversible: %s\n"+
			"Components: %s\n"+
			"Special Components: %s\n"+
			"Class: %s\n"+
			"Description:\n\n%s\n",
		s.SpellName,
		s.Level,
		s.Frequency,
		s.Range,
		s.CastingTime,
		s.Duration,
		s.AreaOfEffect,
		s.Damage,
		s.SavingThrow,
		s.Reversible,
		s.Components,
		s.SpecialComponents,
		s.Class,
		s.formatDescription(),
	)

	return sb
}

func (s *ArcanistSpell) formatDescription() string {
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
