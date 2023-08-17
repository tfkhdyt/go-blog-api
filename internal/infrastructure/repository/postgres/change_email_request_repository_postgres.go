package postgres

import (
	"gorm.io/gorm"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type ChangeEmailRequestRepositoryPostgres struct {
	db *gorm.DB `di.inject:"database"`
}

func (r *ChangeEmailRequestRepositoryPostgres) AddRequest(
	user *entity.User,
	request *entity.ChangeEmailRequest,
) (*entity.ChangeEmailRequest, error) {
	if err := r.db.
		Model(user).
		Association("ChangeEmailRequests").
		Append(request); err != nil {
		return nil, exception.
			NewHTTPError(500, "failed to add change-email request")
	}

	return request, nil
}

func (r *ChangeEmailRequestRepositoryPostgres) RemoveRequestByToken(
	token string,
) error {
	if err := r.db.Delete(
		&entity.ChangeEmailRequest{},
		"token = ?",
		token,
	).Error; err != nil {
		return exception.NewHTTPError(500, "failed to delete change-email request")
	}

	return nil
}

func (r *ChangeEmailRequestRepositoryPostgres) FindRequestByToken(
	token string,
) (*entity.ChangeEmailRequest, error) {
	request := new(entity.ChangeEmailRequest)
	if err := r.db.First(request, "token = ?", token).Error; err != nil {
		return nil, exception.NewHTTPError(404, "request is not found")
	}

	return request, nil
}
