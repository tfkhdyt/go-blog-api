package database

import (
	"errors"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"codeberg.org/tfkhdyt/blog-api/config"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/security"
)

var (
	PostgresInstance *gorm.DB
	err              error
)

func seedAdmin() {
	bcryptService := security.BcryptService{}

	admin := &entity.User{
		FullName: "admin",
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: os.Getenv("ADMIN_PASSWORD"),
		Role:     "admin",
		Username: "admin",
	}

	hashedPassword, err := bcryptService.HashPassword(admin.Password)
	if err != nil {
		log.Fatalln("Error:", err.Error())
	}

	admin.Password = hashedPassword

	if err := PostgresInstance.Create(admin).Error; err != nil {
		log.Fatalln("Error:", err.Error())
	}

	log.Println("Admin account seed success!")
}

func initPostgres() {
	PostgresInstance, err = gorm.
		Open(postgres.Open(config.PostgesDSN), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error:", err.Error())
	}

	if err := PostgresInstance.AutoMigrate(
		&entity.User{},
		&entity.Auth{},
		&entity.ResetPasswordToken{},
		&entity.ChangeEmailRequest{},
	); err != nil {
		log.Fatalln("Error:", err.Error())
	}

	if PostgresInstance.Migrator().HasTable(&entity.User{}) {
		if err := PostgresInstance.
			First(
				&entity.User{},
				"role = ?",
				"admin",
			).
			Error; errors.Is(err, gorm.ErrRecordNotFound) {
			seedAdmin()
		}
	}

	log.Println("Connected to DB...")
}
