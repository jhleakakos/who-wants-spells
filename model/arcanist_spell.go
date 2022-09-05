package model

import (
	"encoding/csv"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
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
		"Spell Name: %s\n"+
			"Level: %s\n"+
			"Frequency: %s\n"+
			"Description: %s\n"+
			"Range: %s\n"+
			"Casting Time: %s\n"+
			"Duration: %s\n"+
			"Area of Effect: %s\n"+
			"Damage: %s\n"+
			"Saving Throw: %s\n"+
			"Reversible: %s\n"+
			"Components: %s\n"+
			"Special Components: %s\n"+
			"Class: %s\n",
		s.SpellName,
		s.Level,
		s.Frequency,
		s.Description,
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
	)

	return sb
}

func formatDescription(s string) string {
	return s
}

func LoadArcanistDataFromCSV(file string) *map[string]ArcanistSpell {
	infile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()

	data := csv.NewReader(infile)
	arcanistSpellBook := make(map[string]ArcanistSpell)

	// skip header row
	data.Read()

	for {
		row, err := data.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		spell, err := convertArcanistCSVRowToStruct(row)
		if err != nil {
			log.Print(err)
		}
		arcanistSpellBook[row[1]] = *spell
	}

	return &arcanistSpellBook
}

func convertArcanistCSVRowToStruct(row []string) (*ArcanistSpell, error) {

	if len(row) != 14 {
		return nil, errors.New("csv row must have 14 fields")
	}

	rowStruct := ArcanistSpell{
		Level:             row[0],
		SpellName:         row[1],
		Frequency:         row[2],
		Description:       row[3],
		Range:             row[4],
		CastingTime:       row[5],
		Duration:          row[6],
		AreaOfEffect:      row[7],
		Damage:            row[8],
		SavingThrow:       row[9],
		Reversible:        row[10],
		Components:        row[11],
		SpecialComponents: row[12],
		Class:             row[13],
	}

	return &rowStruct, nil
}
