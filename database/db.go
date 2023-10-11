package database

import (
	"fmt"

	"github.com/ecommerce_cart/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "postgresql://localhost:5432/postgres?user=shivam&password=shivam"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Println("failed to connect database", err.Error())
		return
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
