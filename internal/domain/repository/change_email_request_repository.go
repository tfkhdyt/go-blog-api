package repository

import "codeberg.org/tfkhdyt/blog-api/internal/domain/entity"

type ChangeEmailRequestRepository interface {
	AddRequest(
		user *entity.User,
		request *entity.ChangeEmailRequest,
	) (*entity.ChangeEmailRequest, error)
	FindRequestByToken(token string) (*entity.ChangeEmailRequest, error)
	RemoveRequestByToken(token string) error
}
