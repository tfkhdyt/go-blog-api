package usecase

import (
	"fmt"
	"time"

	"codeberg.org/tfkhdyt/blog-api/internal/application/dto"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/repository"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/service"
)

type ResetPasswordTokenUsecase struct {
	resetPasswordTokenRepo repository.ResetPasswordTokenRepository `di.inject:"resetPasswordTokenRepo"`
	userRepo               repository.UserRepository               `di.inject:"userRepo"`
	idService              service.IDService                       `di.inject:"idService"`
}

func (r *ResetPasswordTokenUsecase) GetResetPasswordToken(
	payload *dto.GetResetPasswordTokenRequest,
) (*dto.GetResetPasswordTokenResponse, error) {
	user, err := r.userRepo.FindOneUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	token := r.idService.GenerateID()

	result, errAddToken := r.resetPasswordTokenRepo.AddToken(user, &entity.ResetPasswordToken{
		Token:     token,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	})
	if errAddToken != nil {
		return nil, errAddToken
	}

	response := dto.GetResetPasswordTokenResponse{
		Message: fmt.Sprintf("your token is %s", result.Token),
	}

	return &response, nil
}
