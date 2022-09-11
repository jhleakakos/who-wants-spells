package model

import (
	"fmt"
	"gorm.io/gorm"
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
		FormatDescription(s.Description),
	)

	return sb
}
