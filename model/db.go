package model

import (
	"fmt"
	"gorm.io/gorm"
)

const arcanistFile = "./data/arcanist.csv"

func InitializeDB(db *gorm.DB) {

	db.AutoMigrate(&ArcanistSpell{})

	arcanistSpellBook := LoadArcanistDataFromCSV(arcanistFile)
	for key, spell := range *arcanistSpellBook {
		fmt.Printf("loading %s\n", key)
		db.Create(&spell)
	}
}
