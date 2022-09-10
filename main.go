package main

import (
	"fmt"
	"github.com/jhleakakos/who-wants-spells/controller"
	"github.com/jhleakakos/who-wants-spells/view"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func init() {
	var err error
	controller.DB, err = gorm.Open(sqlite.Open("./spells.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection no good")
	}
}

func main() {
	fmt.Println(view.PrintWelcomeScreen())

	// comment or uncomment this as needed to rebuild db
	//controller.InitializeDB()

	controller.RunMenuLoop()

}
