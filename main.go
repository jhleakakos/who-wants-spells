package main

import (
	"fmt"
	"github.com/jhleakakos/who-wants-spells/controller"
	"github.com/jhleakakos/who-wants-spells/view"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

func init() {
	var err error

	// this sets the db location in the same directory as
	// the executable that runs the program
	//
	// this means you just need to put the db in the same
	// directory as the executable when distributing the
	// app to other systems
	executable, err := os.Executable()
	if err != nil {
		log.Fatal("problem getting executable directory")
	}
	controller.ExecutableDir = filepath.Dir(executable)
	dbLocation := fmt.Sprint(controller.ExecutableDir, "/spells.sqlite")

	controller.DB, err = gorm.Open(sqlite.Open(dbLocation), &gorm.Config{})
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
