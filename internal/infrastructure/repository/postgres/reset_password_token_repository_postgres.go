package postgres

import (
	"gorm.io/gorm"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type ResetPasswordTokenRepositoryPostgres struct {
	db *gorm.DB `di.inject:"database"`
}

func (r *ResetPasswordTokenRepositoryPostgres) AddToken(
	user *entity.User,
	token *entity.ResetPasswordToken,
) (*entity.ResetPasswordToken, error) {
	if err := r.db.
		Model(user).
		Association("ResetPasswordTokens").
		Append(token); err != nil {
		return nil, exception.
			NewHTTPError(500, "failed to add reset password token")
	}

	return token, nil
}

func (r *ResetPasswordTokenRepositoryPostgres) RemoveToken(
	token string,
) error {
	if err := r.db.Delete(
		&entity.ResetPasswordToken{},
		"token = ?",
		token,
	).Error; err != nil {
		return exception.NewHTTPError(500, "failed to delete reset password token")
	}

	return nil
}

func (r *ResetPasswordTokenRepositoryPostgres) FindToken(
	token string,
) (*entity.ResetPasswordToken, error) {
	tkn := new(entity.ResetPasswordToken)
	if err := r.db.First(tkn, "token = ?", token).Error; err != nil {
		return nil, exception.NewHTTPError(404, "token is not found")
	}

	return tkn, nil
}
