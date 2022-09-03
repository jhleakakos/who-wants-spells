package model

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

const arcanistFile = "./data/arcanist.csv"

func InitializeDB(db *gorm.DB) {

	err := db.Migrator().DropTable(&ArcanistSpell{})
	if err != nil {
		log.Print("check DropTable()")
	}

	err = db.AutoMigrate(&ArcanistSpell{})
	if err != nil {
		log.Print("check AutoMigrate()")
	}

	arcanistSpellBook := LoadArcanistDataFromCSV(arcanistFile)
	for key, spell := range *arcanistSpellBook {
		fmt.Printf("loading %s\n", key)
		db.Create(&spell)
	}
}
