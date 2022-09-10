package controller

import (
	"fmt"
	"github.com/jhleakakos/who-wants-spells/model"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

const arcanistFile = "../data/arcanist.csv"

func InitializeDB() {

	err := DB.Migrator().DropTable(&model.ArcanistSpell{})
	if err != nil {
		log.Print("check DropTable()")
	}

	err = DB.AutoMigrate(&model.ArcanistSpell{})
	if err != nil {
		log.Print("check AutoMigrate()")
	}

	arcanistSpellBook := LoadArcanistDataFromCSV(arcanistFile)
	for key, spell := range *arcanistSpellBook {
		fmt.Printf("loading %s\n", key)
		DB.Create(&spell)
	}
}
