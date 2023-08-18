package postgres

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"codeberg.org/tfkhdyt/blog-api/config"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/database/postgres/sqlc"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/security"
)

func seedAdmin(ctx context.Context, db sqlc.Querier) {
	bcryptService := security.BcryptService{}

	admin := &entity.User{
		FullName: "admin",
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: os.Getenv("ADMIN_PASSWORD"),
		Role:     entity.RoleAdmin,
		Username: "admin",
	}

	var errHash error
	admin.Password, errHash = bcryptService.HashPassword(admin.Password)
	if errHash != nil {
		log.Fatalln("ERROR:", errHash)
	}

	admins, err := db.FindAdmin(ctx)
	if err != nil {
		log.Fatalln("ERROR: failed to find admin")
	}
	if len(admins) == 0 {
		if _, err := db.CreateUser(
			ctx,
			sqlc.CreateUserParams{
				FullName: admin.FullName,
				Email:    admin.Email,
				Password: admin.Password,
				Role:     sqlc.NullRole{Role: sqlc.Role(admin.Role), Valid: true},
				Username: admin.Username,
			},
		); err != nil {
			log.Fatalln("ERROR: failed to seed admin")
		} else {
			log.Println("INFO: Admin account seed success!")
		}
	}
}

func GetPostgresSQLCQuerier(ctx context.Context) *sqlc.Queries {
	conn, err := pgxpool.New(ctx, config.PostgresURL)
	if err != nil {
		log.Fatalln("ERROR:", err.Error())
	}

	db := sqlc.New(conn)
	seedAdmin(ctx, db)

	log.Println("INFO:", "Connected to DB")

	return db
}
