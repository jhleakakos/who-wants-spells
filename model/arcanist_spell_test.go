package model

import (
	"testing"
)

func TestConvertCSVRowToStruct(t *testing.T) {

	inputValid := []string{
		"1",
		"chain weevil",
		"super duper rare",
		"They eat all your skins, and they never apologize",
		"5 ft",
		"instantaneous",
		"10 rounds",
		"weevil ball",
		"2d4 hp per round",
		"roll reaction checks",
		"no",
		"bugs",
		"more bugs",
		"arcanist",
	}

	got, err := convertArcanistCSVRowToStruct(inputValid)
	if err != nil {
		t.Errorf("throwing an unexpected error: %s", err)
	}

	expected := ArcanistSpell{
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

	if *got != expected {
		t.Error("problem converting csv row to struct for arcanist")
	}

	inputInvalidTooFewColumns := []string{
		"1",
		"chain weevil",
		"super duper rare",
		"They eat all your skins, and they never apologize",
		"5 ft",
		"instantaneous",
		"10 rounds",
		"weevil ball",
		"2d4 hp per round",
		"roll reaction checks",
		"no",
		"bugs",
		"more bugs",
	}

	_, err = convertArcanistCSVRowToStruct(inputInvalidTooFewColumns)
	if err == nil {
		t.Error("expecting error for too few columns")
	}

	inputInvalidTooManyColumns := []string{
		"1",
		"chain weevil",
		"super duper rare",
		"They eat all your skins, and they never apologize",
		"5 ft",
		"instantaneous",
		"10 rounds",
		"weevil ball",
		"2d4 hp per round",
		"roll reaction checks",
		"no",
		"bugs",
		"more bugs",
		"arcanist",
		"periwinkle",
	}

	_, err = convertArcanistCSVRowToStruct(inputInvalidTooManyColumns)
	if err == nil {
		t.Error("expecting error for too many columns")
	}

}
