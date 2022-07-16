package view

import "fmt"

func DisplayMenu() string {
	return fmt.Sprint("Select choice by number:\n\n" +
		"1- Arcanist Spell Book\n" +
		"2- Diviner Spell Book\n\n" +
		"10- Look up spell by name\n" +
		"0- Quit\n\n" +
		"Select choice:")
}
