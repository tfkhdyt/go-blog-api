package repository

import "codeberg.org/tfkhdyt/blog-api/internal/domain/entity"

type ResetPasswordTokenRepository interface {
	AddToken(
		user *entity.User,
		token *entity.ResetPasswordToken,
	) (*entity.ResetPasswordToken, error)
	FindToken(token string) (*entity.ResetPasswordToken, error)
	RemoveToken(token string) error
}
