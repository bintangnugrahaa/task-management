package config

import (
	"fmt"
	"log"
	"tusk/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbHost     = "localhost"
	dbPort     = 3306
	dbUser     = "root"
	dbPassword = "root"
	dbName     = "tusk"
)

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return db
}

func CreateOwnerAccount(db *gorm.DB) {
	const (
		defaultPassword = "123456"
		defaultEmail    = "owner@go.id"
	)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
	}

	owner := models.User{
		Role:     "Admin",
		Name:     "Owner",
		Email:    defaultEmail,
		Password: string(hashedPassword),
	}

	var existing models.User
	if err := db.Where("email = ?", owner.Email).First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&owner).Error; err != nil {
				log.Printf("failed to create owner account: %v", err)
			} else {
				log.Println("Owner account created.")
			}
		} else {
			log.Printf("failed to query owner account: %v", err)
		}
	} else {
		log.Println("Owner account already exists.")
	}
}
