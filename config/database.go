package config

import (
	"fmt"
	"tusk/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = ""
	dbName   = "tusk"
)

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, port, dbName,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	return database
}

func CreateOwnerAccount(db *gorm.DB) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		panic("Failed to hash password: " + err.Error())
	}

	owner := models.User{
		Role:     "Admin",
		Name:     "Owner",
		Email:    "owner@go.id",
		Password: string(hashedPassword),
	}

	var existing models.User
	if db.Where("email = ?", owner.Email).First(&existing).RowsAffected == 0 {
		db.Create(&owner)
		fmt.Println("Owner account created.")
	} else {
		fmt.Println("Owner already exists.")
	}
}
