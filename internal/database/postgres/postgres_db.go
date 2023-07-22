package postgres

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	postgresConfig "codeberg.org/tfkhdyt/blog-api/internal/config/postgres"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/user"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	DB, err = gorm.Open(postgres.Open(postgresConfig.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error:", err.Error())
	}

	if err := DB.AutoMigrate(&user.User{}); err != nil {
		log.Fatalln("Error:", err.Error())
	}

	log.Println("Connected to DB...")
}
