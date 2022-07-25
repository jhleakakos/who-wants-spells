package main

import (
	"fmt"
	"github.com/jhleakakos/who-wants-spells/model"
	"github.com/jhleakakos/who-wants-spells/view"
)

const file = "./data/arcanist.csv"

func main() {
	fmt.Println(view.DisplayMenu())
	fmt.Println()
	arcanistSpellBook := model.LoadData(file)
	for key, spell := range *arcanistSpellBook {
		fmt.Printf("%s: %s\n\n\n\n", key, spell)
	}
}
