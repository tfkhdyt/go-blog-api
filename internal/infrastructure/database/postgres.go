package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/security"
)

var (
	DB         *gorm.DB
	err        error
	dbUsername = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbName     = os.Getenv("DB_NAME")
	dsn        = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUsername, dbPassword, dbName, dbPort,
	)
)

func seedAdmin() {
	bcryptService := security.BcryptService{}

	admin := &entity.User{
		FullName: "admin",
		Email:    "admin@tfkhdyt.my.id",
		Password: os.Getenv("ADMIN_PASSWORD"),
		Role:     "admin",
		Username: "admin",
	}

	hashedPassword, err := bcryptService.HashPassword(admin.Password)
	if err != nil {
		log.Fatalln("Error:", err.Error())
	}

	admin.Password = hashedPassword

	if err := DB.Create(admin).Error; err != nil {
		log.Fatalln("Error:", err.Error())
	}

	log.Println("Admin account seed success!")
}

func init() {
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error:", err.Error())
	}

	if err := DB.AutoMigrate(&entity.User{}, &entity.Auth{}); err != nil {
		log.Fatalln("Error:", err.Error())
	}

	if DB.Migrator().HasTable(&entity.User{}) {
		if err := DB.First(&entity.User{}, "role = ?", "admin").Error; errors.Is(
			err,
			gorm.ErrRecordNotFound,
		) {
			seedAdmin()
		}
	}

	log.Println("Connected to DB...")
}
