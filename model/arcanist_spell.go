package model

import "errors"

func LoadData() {
	// csv.NewReader() accepts io.Reader
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
