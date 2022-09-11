package controller

import (
	"fmt"
	"github.com/jhleakakos/who-wants-spells/model"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var ExecutableDir string

const arcanistFile = "./data/arcanist.csv"
const divinerFile = "./data/diviner.csv"

func InitializeDB() {
	var err error
	err = DB.Migrator().DropTable(&model.ArcanistSpell{})
	if err != nil {
		log.Print("check arcanist DropTable()")
	}

	err = DB.Migrator().DropTable(&model.DivinerSpell{})
	if err != nil {
		log.Print("check diviner DropTable()")
	}

	err = DB.AutoMigrate(&model.ArcanistSpell{})
	if err != nil {
		log.Print("check arcanist AutoMigrate()")
	}

	err = DB.AutoMigrate(&model.DivinerSpell{})
	if err != nil {
		log.Print("check diviner AutoMigrate()")
	}

	arcanistSpellBook := LoadArcanistDataFromCSV(arcanistFile)
	fmt.Println("\n\nLoading arcanist spells\n")
	for key, spell := range *arcanistSpellBook {
		fmt.Printf("loading %s\n", key)
		DB.Create(&spell)
	}

	divinerSpellBook := LoadDivinerDataFromCSV(divinerFile)
	fmt.Println("\n\nLoading diviner spells\n")
	for key, spell := range *divinerSpellBook {
		fmt.Printf("loading %s\n", key)
		DB.Create(&spell)
	}
}
