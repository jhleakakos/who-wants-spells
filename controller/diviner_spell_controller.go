package controller

import (
	"encoding/csv"
	"errors"
	"github.com/jhleakakos/who-wants-spells/model"
	"io"
	"log"
	"os"
)

func LoadDivinerDataFromCSV(file string) *map[string]model.DivinerSpell {
	infile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()

	data := csv.NewReader(infile)
	divinerSpellBook := make(map[string]model.DivinerSpell)

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

		spell, err := convertDivinerCSVRowToStruct(row)
		if err != nil {
			log.Print(err)
		}
		divinerSpellBook[row[2]] = *spell
	}

	return &divinerSpellBook
}

func convertDivinerCSVRowToStruct(row []string) (*model.DivinerSpell, error) {
	if len(row) != 11 {
		return nil, errors.New("diviner csv row must have 11 fields")
	}

	rowStruct := model.DivinerSpell{
		DivinationSchool: row[0],
		Tier:             row[1],
		Name:             row[2],
		Description:      row[3],
		Range:            row[4],
		CastingTime:      row[5],
		Duration:         row[6],
		AreaOfEffect:     row[7],
		Damage:           row[8],
		SavingThrow:      row[9],
		Reversible:       row[10],
	}

	return &rowStruct, nil
}
