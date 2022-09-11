package model

import (
	"fmt"
	"gorm.io/gorm"
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
		FormatDescription(s.Description),
	)

	return sb
}
