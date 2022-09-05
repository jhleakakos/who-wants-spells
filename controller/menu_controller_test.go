package controller

import "testing"

func TestDisplayMenu(t *testing.T) {

	got := DisplayMenu()

	expected := "Select choice by number:\n\n" +
		"1- Arcanist Spell Book\n" +
		"2- Diviner Spell Book\n\n" +
		"10- Look up spell by name\n" +
		"0- Quit\n\n" +
		"Select choice:"

	if got != expected {
		t.Error("Check your menu")
	}

}
