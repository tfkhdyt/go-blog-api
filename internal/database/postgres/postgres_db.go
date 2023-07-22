package postgres

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	postgresConfig "codeberg.org/tfkhdyt/blog-api/internal/config/postgres"
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

	if err := DB.AutoMigrate(); err != nil {
		log.Fatalln("Error:", err.Error())
	}

	log.Println("Connected to DB...")
}
