package main

import (
	"fmt"
	"github.com/jhleakakos/who-wants-spells/model"
	"github.com/jhleakakos/who-wants-spells/view"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

const file = "./data/arcanist.csv"

func main() {
	fmt.Println(view.DisplayMenu())
	fmt.Println()

	db, err := gorm.Open(sqlite.Open("./data/spells.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection no good")
	}

	db.AutoMigrate(&model.ArcanistSpell{})

	arcanistSpellBook := model.LoadArcanistDataFromCSV(file)
	for key, spell := range *arcanistSpellBook {
		fmt.Printf("loading %s\n", key)
		db.Create(&spell)
	}

	var arcanistSpell model.ArcanistSpell
	db.First(&arcanistSpell)
	fmt.Println(arcanistSpell)
}
