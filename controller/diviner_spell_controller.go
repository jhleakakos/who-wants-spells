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

var divinerSpells []model.DivinerSpell

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

func displayDivinerSpells() {
	for _, spell := range divinerSpells {
		fmt.Println(spell.PrintSpell())
	}
}

func getDivinerSpellByDivinationSchool(spellDivinationSchool string) {
	DB.Where("divination_school like ?", fmt.Sprint("%", spellDivinationSchool, "%")).Find(&divinerSpells)
	displayDivinerSpells()
}

func getDivinerSpellByTier(spellTier string) {
	DB.Where("tier like ?", fmt.Sprint("%", spellTier, "%")).Find(&divinerSpells)
	displayDivinerSpells()
}

func getDivinerSpellByName(spellName string) {
	DB.Where("name like ?", fmt.Sprint("%", spellName, "%")).Find(&divinerSpells)
	displayDivinerSpells()
}

func getDivinerSpellByDescription(spellDescription string) {
	DB.Where("description like ?", fmt.Sprint("%", spellDescription, "%")).Find(&divinerSpells)
	displayDivinerSpells()
}

func getDivinerSpellByRange(spellRange string) {
	DB.Where("range like ?", fmt.Sprint("%", spellRange, "%")).Find(&divinerSpells)
	displayDivinerSpells()
}

func getDivinerSpellByCastingTime(spellCastingTime string) {
	DB.Where("casting_time like ?", fmt.Sprint("%", spellCastingTime, "%")).Find(&divinerSpells)
	displayDivinerSpells()
}

func getDivinerSpellByDuration(spellDuration string) {
	DB.Where("duration like ?", fmt.Sprint("%", spellDuration, "%")).Find(&divinerSpells)
	displayDivinerSpells()
}

func getDivinerSpellByAreaOfEffect(spellAreaOfEffect string) {
	DB.Where("area_of_effect like ?", fmt.Sprint("%", spellAreaOfEffect, "%")).Find(&divinerSpells)
	displayDivinerSpells()
}

func getDivinerSpellByDamage(spellDamage string) {
	DB.Where("damage like ?", fmt.Sprint("%", spellDamage, "%")).Find(&divinerSpells)
	displayDivinerSpells()
}

func getDivinerSpellBySavingThrow(spellSavingThrow string) {
	DB.Where("saving_throw like ?", fmt.Sprint("%", spellSavingThrow, "%")).Find(&divinerSpells)
	displayDivinerSpells()
}

func getDivinerSpellByReversible(spellReversible string) {
	var searchVal string
	if spellReversible == "y" {
		searchVal = "Yes"
	}
	if spellReversible == "n" {
		searchVal = "No"
	}

	DB.Where("reversible = ?", searchVal).Find(&divinerSpells)
	displayDivinerSpells()
}

func getAllDivinerSpells() {
	DB.Find(&divinerSpells)
	displayDivinerSpells()
}
