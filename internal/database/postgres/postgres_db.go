package postgres

import (
	"errors"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	postgresConfig "codeberg.org/tfkhdyt/blog-api/internal/config/postgres"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/auth"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/user"
)

var (
	DB  *gorm.DB
	err error
)

func seedAdmin() {
	admin := &user.User{
		FullName: "admin",
		Email:    "admin@tfkhdyt.my.id",
		Password: os.Getenv("ADMIN_PASSWORD"),
		Role:     "admin",
		Username: "admin",
	}

	if err := admin.HashPassword(); err != nil {
		log.Fatalln("Error:", err.Error())
	}

	if err := DB.Create(admin).Error; err != nil {
		log.Fatalln("Error:", err.Error())
	}

	log.Println("Admin account seed success!")
}

func init() {
	DB, err = gorm.Open(postgres.Open(postgresConfig.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error:", err.Error())
	}

	if err := DB.AutoMigrate(&user.User{}, &auth.Auth{}); err != nil {
		log.Fatalln("Error:", err.Error())
	}

	if DB.Migrator().HasTable(&user.User{}) {
		if err := DB.First(&user.User{}, "role = ?", "admin").Error; errors.Is(
			err,
			gorm.ErrRecordNotFound,
		) {
			seedAdmin()
		}
	}

	log.Println("Connected to DB...")
}
