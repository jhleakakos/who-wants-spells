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

func getArcanistSpellByName(spellName string) {
	var arcanistSpells []model.ArcanistSpell
	DB.Where("spell_name like ?", fmt.Sprint("%", spellName, "%")).Find(&arcanistSpells)
	displayArcanistSpells(arcanistSpells...)
}

func getAllArcanistSpells() {
	var arcanistSpells []model.ArcanistSpell
	DB.Find(&arcanistSpells)
	displayArcanistSpells(arcanistSpells...)
}
