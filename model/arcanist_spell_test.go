package model

import (
	"testing"
)

func TestArcanistSpellPrintSpell(t *testing.T) {
	testData := ArcanistSpell{
		Level:             "1",
		SpellName:         "chain weevil",
		Frequency:         "super duper rare",
		Description:       "They eat all your skins, and they never apologize",
		Range:             "5 ft",
		CastingTime:       "instantaneous",
		Duration:          "10 rounds",
		AreaOfEffect:      "weevil ball",
		Damage:            "2d4 hp per round",
		SavingThrow:       "roll reaction checks",
		Reversible:        "no",
		Components:        "bugs",
		SpecialComponents: "more bugs",
		Class:             "arcanist",
	}

	inputValid := "" +
		"Spell Name: chain weevil\n" +
		"Level: 1\n" +
		"Frequency: super duper rare\n" +
		"Description: They eat all your skins, and they never apologize\n" +
		"Range: 5 ft\n" +
		"Casting Time: instantaneous\n" +
		"Duration: 10 rounds\n" +
		"Area of Effect: weevil ball\n" +
		"Damage: 2d4 hp per round\n" +
		"Saving Throw: roll reaction checks\n" +
		"Reversible: no\n" +
		"Components: bugs\n" +
		"Special Components: more bugs\n" +
		"Class: arcanist\n"

	got := testData.PrintSpell()
	if got != inputValid {
		t.Error("arcanist spell print output incorrect")
	}
}
