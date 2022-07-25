package model

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
)

func LoadData(file string) *map[string]ArcanistSpell {
	infile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()

	data := csv.NewReader(infile)
	arcanistSpellBook := make(map[string]ArcanistSpell)

	for {
		row, err := data.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		spell, err := ConvertCSVRowToStruct(row)
		if err != nil {
			log.Print(err)
		}
		arcanistSpellBook[row[1]] = *spell
	}

	return &arcanistSpellBook
}

func ConvertCSVRowToStruct(row []string) (*ArcanistSpell, error) {

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
