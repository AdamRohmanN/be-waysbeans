package database

import (
	"fmt"
	"waysbeans/models"
	"waysbeans/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Product{},
		&models.Category{},
		&models.Transaction{},
	)

	if err != nil {
		fmt.Println(err)
		panic("migration failed")
	}

	fmt.Println("migration success")
}
