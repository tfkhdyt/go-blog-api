package usecase

import (
	"fmt"
	"time"

	"codeberg.org/tfkhdyt/blog-api/config"
	"codeberg.org/tfkhdyt/blog-api/internal/application/dto"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/repository"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/service"
)

type ResetPasswordTokenUsecase struct {
	resetPasswordTokenRepo repository.ResetPasswordTokenRepository `di.inject:"resetPasswordTokenRepo"`
	userRepo               repository.UserRepository               `di.inject:"userRepo"`
	idService              service.IDService                       `di.inject:"idService"`
	emailService           service.EmailService                    `di.inject:"emailService"`
}

func (r *ResetPasswordTokenUsecase) GetResetPasswordToken(
	payload *dto.GetResetPasswordTokenRequest,
) (*dto.GetResetPasswordTokenResponse, error) {
	user, err := r.userRepo.FindOneUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	token := r.idService.GenerateID()

	if _, errAddToken := r.resetPasswordTokenRepo.AddToken(
		user,
		&entity.ResetPasswordToken{
			Token:     token,
			ExpiresAt: time.Now().Add(5 * time.Minute),
		},
	); errAddToken != nil {
		return nil, errAddToken
	}

	if err := r.emailService.SendMail(
		entity.Recipient{
			Email: config.MailSenderEmail,
			Name:  config.MailSenderName,
		}, entity.Recipient{
			Email: user.Email,
			Name:  user.FullName,
		},
		"Reset Password Token",
		fmt.Sprintf("Your reset password token is <b>%v<b>", token),
	); err != nil {
		return nil, err
	}

	response := dto.GetResetPasswordTokenResponse{
		Message: "the password reset token has been sent to your email.",
	}

	return &response, nil
}
