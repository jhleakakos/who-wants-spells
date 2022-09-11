package controller

import (
	"encoding/csv"
	"errors"
	"fmt"
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

func displayDivinerSpells(spells ...model.DivinerSpell) {
	for _, spell := range spells {
		fmt.Println(spell.PrintSpell())
	}
}

func getDivinerSpellByDivinationSchool(spellDivinationSchool string) {
	var divinerSpells []model.DivinerSpell
	DB.Where("divination_school like ?", fmt.Sprint("%", spellDivinationSchool, "%")).Find(&divinerSpells)
	displayDivinerSpells(divinerSpells...)
}

func getDivinerSpellByTier(spellTier string) {
	var divinerSpells []model.DivinerSpell
	DB.Where("tier like ?", fmt.Sprint("%", spellTier, "%")).Find(&divinerSpells)
	displayDivinerSpells(divinerSpells...)
}

func getDivinerSpellByName(spellName string) {
	var divinerSpells []model.DivinerSpell
	DB.Where("name like ?", fmt.Sprint("%", spellName, "%")).Find(&divinerSpells)
	displayDivinerSpells(divinerSpells...)
}

func getDivinerSpellByDescription(spellDescription string) {
	var divinerSpells []model.DivinerSpell
	DB.Where("description like ?", fmt.Sprint("%", spellDescription, "%")).Find(&divinerSpells)
	displayDivinerSpells(divinerSpells...)
}

func getDivinerSpellByRange(spellRange string) {
	var divinerSpells []model.DivinerSpell
	DB.Where("range like ?", fmt.Sprint("%", spellRange, "%")).Find(&divinerSpells)
	displayDivinerSpells(divinerSpells...)
}

func getDivinerSpellByCastingTime(spellCastingTime string) {
	var divinerSpells []model.DivinerSpell
	DB.Where("casting_time like ?", fmt.Sprint("%", spellCastingTime, "%")).Find(&divinerSpells)
	displayDivinerSpells(divinerSpells...)
}

func getDivinerSpellByDuration(spellDuration string) {
	var divinerSpells []model.DivinerSpell
	DB.Where("duration like ?", fmt.Sprint("%", spellDuration, "%")).Find(&divinerSpells)
	displayDivinerSpells(divinerSpells...)
}

func getDivinerSpellByAreaOfEffect(spellAreaOfEffect string) {
	var divinerSpells []model.DivinerSpell
	DB.Where("area_of_effect like ?", fmt.Sprint("%", spellAreaOfEffect, "%")).Find(&divinerSpells)
	displayDivinerSpells(divinerSpells...)
}

func getDivinerSpellByDamage(spellDamage string) {
	var divinerSpells []model.DivinerSpell
	DB.Where("damage like ?", fmt.Sprint("%", spellDamage, "%")).Find(&divinerSpells)
	displayDivinerSpells(divinerSpells...)
}
