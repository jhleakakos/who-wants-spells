package controller

import (
	"fmt"
	"github.com/jhleakakos/who-wants-spells/model"
)

func displayArcanistSpells(spells ...model.ArcanistSpell) {
	for _, spell := range spells {
		fmt.Println(spell.PrintSpell())
	}
}

func getArcanistSpellByLevel(spellLevel string) {
	var arcanistSpells []model.ArcanistSpell
	DB.Where("level = ?", spellLevel).Find(&arcanistSpells)
	displayArcanistSpells(arcanistSpells...)
}

func getArcanistSpellByName(spellName string) {
	var arcanistSpells []model.ArcanistSpell
	DB.Where("spell_name like ?", fmt.Sprint("%", spellName, "%")).Find(&arcanistSpells)
	displayArcanistSpells(arcanistSpells...)
}

func getArcanistSpellByFrequency(spellFrequency string) {
	var arcanistSpells []model.ArcanistSpell
	DB.Where("frequency like ?", fmt.Sprint("%", spellFrequency, "%")).Find(&arcanistSpells)
	displayArcanistSpells(arcanistSpells...)
}

func getArcanistSpellByDescription(spellDescription string) {
	var arcanistSpells []model.ArcanistSpell
	DB.Where("description like ?", fmt.Sprint("%", spellDescription, "%")).Find(&arcanistSpells)
	displayArcanistSpells(arcanistSpells...)
}

func getArcanistSpellByRange(spellRange string) {
	var arcanistSpells []model.ArcanistSpell
	DB.Where("range like ?", fmt.Sprint("%", spellRange, "%")).Find(&arcanistSpells)
	displayArcanistSpells(arcanistSpells...)
}

func getArcanistSpellByCastingTime(spellCastingTime string) {
	var arcanistSpells []model.ArcanistSpell
	DB.Where("casting_time like ?", fmt.Sprint("%", spellCastingTime, "%")).Find(&arcanistSpells)
	displayArcanistSpells(arcanistSpells...)
}

func getArcanistSpellByDuration(spellDuration string) {
	var arcanistSpells []model.ArcanistSpell
	DB.Where("duration like ?", fmt.Sprint("%", spellDuration, "%")).Find(&arcanistSpells)
	displayArcanistSpells(arcanistSpells...)
}

func getArcanistSpellByAreaOfEffect(spellAreaOfEffect string) {
	var arcanistSpells []model.ArcanistSpell
	DB.Where("area_of_effect like ?", fmt.Sprint("%", spellAreaOfEffect, "%")).Find(&arcanistSpells)
	displayArcanistSpells(arcanistSpells...)
}

func getAllArcanistSpells() {
	var arcanistSpells []model.ArcanistSpell
	DB.Find(&arcanistSpells)
	displayArcanistSpells(arcanistSpells...)
}
