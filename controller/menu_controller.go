package controller

import "fmt"

func DisplayMainMenu() string {
	return "\n\n\n\n\nSelect choice by number:\n\n" +
		"1  -- Arcanist Spell Book\n" +
		"2  -- Diviner Spell Book\n\n" +
		"10 -- Look up spell by name\n" +
		"0  -- Quit\n\n" +
		"Select choice: "
}

func DisplayArcanistMenu() string {
	return "\n\n\n\n\nSelect choice by number:\n\n" +
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
		"20 -- All spells\n" +
		"0  -- Return to main menu\n\n" +
		"Select choice: "
}

func RunMenuLoop() {

	mainMenuInput := -1
	contextMenuInput := -1

	for {
		fmt.Print(DisplayMainMenu())
		fmt.Scanln(&mainMenuInput)

		if mainMenuInput == 0 {
			break
		}

		if mainMenuInput == 1 {
			for {
				fmt.Print(DisplayArcanistMenu())
				fmt.Scanln(&contextMenuInput)

				if contextMenuInput == 1 {
					var spellLevel string
					fmt.Print("Level: ")
					fmt.Scanln(&spellLevel)
					getArcanistSpellByLevel(spellLevel)
				}

				if contextMenuInput == 2 {
					var spellName string
					fmt.Print("Spell name: ")
					fmt.Scanln(&spellName)
					getArcanistSpellByName(spellName)
				}

				if contextMenuInput == 3 {
					var spellFrequency string
					fmt.Print("Frequency: ")
					fmt.Scanln(&spellFrequency)
					getArcanistSpellByFrequency(spellFrequency)
				}

				if contextMenuInput == 4 {
					var spellDescription string
					fmt.Print("Description: ")
					fmt.Scanln(&spellDescription)
					getArcanistSpellByDescription(spellDescription)
				}

				if contextMenuInput == 5 {
					var spellRange string
					fmt.Print("Range: ")
					fmt.Scanln(&spellRange)
					getArcanistSpellByRange(spellRange)
				}

				if contextMenuInput == 6 {
					var spellCastingTime string
					fmt.Print("Casting time: ")
					fmt.Scanln(&spellCastingTime)
					getArcanistSpellByCastingTime(spellCastingTime)
				}

				if contextMenuInput == 7 {
					var spellDuration string
					fmt.Print("Duration: ")
					fmt.Scanln(&spellDuration)
					getArcanistSpellByDuration(spellDuration)
				}

				if contextMenuInput == 8 {
					var spellAreaOfEffect string
					fmt.Print("Area of effect: ")
					fmt.Scanln(&spellAreaOfEffect)
					getArcanistSpellByAreaOfEffect(spellAreaOfEffect)
				}

				if contextMenuInput == 9 {
					var spellDamage string
					fmt.Print("Damage: ")
					fmt.Scanln(&spellDamage)
					getArcanistSpellByDamage(spellDamage)
				}

				if contextMenuInput == 10 {
					var spellSavingThrow string
					fmt.Print("Saving throw: ")
					fmt.Scanln(&spellSavingThrow)
					getArcanistSpellBySavingThrow(spellSavingThrow)
				}

				if contextMenuInput == 11 {
					var spellReversible string
					fmt.Print("Reversible (y/n): ")
					fmt.Scanln(&spellReversible)

					for spellReversible != "y" && spellReversible != "n" {
						fmt.Print("Reversible (y/n): ")
						fmt.Scanln(&spellReversible)
					}

					getArcanistSpellByReversible(spellReversible)
				}

				if contextMenuInput == 12 {
					var spellComponents string
					fmt.Print("Components: ")
					fmt.Scanln(&spellComponents)
					getArcanistSpellByComponents(spellComponents)
				}

				if contextMenuInput == 20 {
					getAllArcanistSpells()
				}

				if contextMenuInput == 0 {
					break
				}
			}
		}

		if mainMenuInput == 2 {
			fmt.Println("yo diviner")
		}
	}

}
