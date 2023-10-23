package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitialMigration() {
	DB, err = gorm.Open(
		sqlite.Open(
			"mediaDatabase.db",
		), &gorm.Config{},
	)

	if err != nil {
		fmt.Println(err.Error())
		panic("error with connection to db")
	}

	DB.AutoMigrate(&Directory{})
	DB.AutoMigrate(&Movies{})
	DB.AutoMigrate(&Shows{})

}
