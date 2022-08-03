package main

import (
	"fmt"
	"github.com/jhleakakos/who-wants-spells/model"
	"github.com/jhleakakos/who-wants-spells/view"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	fmt.Println(view.DisplayMenu())
	fmt.Println()

	db, err := gorm.Open(sqlite.Open("./data/spells.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection no good")
	}

	// comment or uncomment these as needed to rebuild db
	//model.InitializeDB(db)

	//var arcanistSpell model.ArcanistSpell
	//db.First(&arcanistSpell)
	var arcanistSpells []model.ArcanistSpell
	db.Find(&arcanistSpells)
	fmt.Println(arcanistSpells)
}
