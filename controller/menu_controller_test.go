package controller

import (
	"testing"
)

func TestDisplayMainMenu(t *testing.T) {

	got := DisplayMainMenu()

	expected := "\n\n\n\n\nSelect choice by number:\n\n" +
		"1  -- Arcanist Spell Book\n" +
		"2  -- Diviner Spell Book\n\n" +
		"10 -- Look up spell by name\n" +
		"0  -- Quit\n\n" +
		"Select choice: "

	if got != expected {
		t.Error("check main menu")
	}
}

func TestDisplayArcanistMenu(t *testing.T) {
	got := DisplayArcanistMenu()

	expected := "\n\n\n\n\nSelect choice by number:\n\n" +
		"1  -- Level\n" +
		"2  -- Spell name\n" +
		"3  -- Frequency\n" +
		"4  -- Description\n" +
		"5  -- Range\n" +
		"6  -- Casting time\n" +
		"7  -- Duration\n" +
		"8  -- Area of Effect\n" +
		"9  -- Damage\n" +
		"10 -- Saving throw\n" +
		"11 -- Reversible\n" +
		"12 -- Components\n" +
		"13 -- Special components\n" +
		"20 -- All spells" +
		"0  -- Return to main menu\n\n" +
		"Select choice: "

	if got != expected {
		t.Error("check arcanist spell menu")
	}
}
