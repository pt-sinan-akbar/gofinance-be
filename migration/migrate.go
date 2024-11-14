package main

import (
	"fmt"
	"github.com/pt-sinan-akbar/Initializers"
	"github.com/pt-sinan-akbar/Entities"
	"gorm.io/gorm"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables")
	}
 
	initializers.ConnectDB(&config)
}


func main() {
	err := initializers.DB.AutoMigrate(
		&entities.Category{},
		&entities.SubCategory{},
		&entities.Transaction{},
		&entities.User{},
		&entities.UserLog{},
	)

	if err != nil {
		return
	}

	Session := initializers.DB.Session(&gorm.Session{
		PrepareStmt: true,
	})
	if Session != nil {
		fmt.Println("? Migration complete")
	}
}

