package usecase

import (
	"context"
	"fmt"
	"time"

	"codeberg.org/tfkhdyt/blog-api/config"
	"codeberg.org/tfkhdyt/blog-api/internal/application/dto"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/repository"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/service"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type ResetPasswordUsecase struct {
	resetPasswordTokenRepo repository.ResetPasswordTokenRepository `di.inject:"resetPasswordTokenRepo"`
	userRepo               repository.UserRepository               `di.inject:"userRepo"`
	idService              service.IDService                       `di.inject:"idService"`
	emailService           service.EmailService                    `di.inject:"emailService"`
	passwordHashService    service.PasswordHashService             `di.inject:"passwordHashService"`
}

func (r *ResetPasswordUsecase) GetResetPasswordToken(
	payload *dto.GetResetPasswordTokenRequest,
) (*dto.GetResetPasswordTokenResponse, error) {
	ctx := context.Background()

	user, err := r.userRepo.FindOneUserByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}

	token := r.idService.GenerateID()

	if _, errAddToken := r.resetPasswordTokenRepo.AddToken(
		ctx,
		user.ID,
		&entity.ResetPasswordToken{
			Token:     token,
			ExpiresAt: time.Now().Add(5 * time.Minute),
		},
	); errAddToken != nil {
		return nil, errAddToken
	}

	if err := r.emailService.SendMail(
		&entity.Recipient{
			Email: config.MailSenderEmail,
			Name:  config.MailSenderName,
		},
		&entity.Recipient{
			Email: user.Email,
			Name:  user.FullName,
		},
		"Reset-Password Token",
		fmt.Sprintf("Your reset-password token is <b>%v<b>", token),
	); err != nil {
		return nil, err
	}

	response := dto.GetResetPasswordTokenResponse{
		Message: "reset-password token has been sent to your email",
	}

	return &response, nil
}

func (r *ResetPasswordUsecase) ResetPassword(
	token string,
	payload *dto.ResetPasswordRequest,
) (*dto.ResetPasswordResponse, error) {
	ctx := context.Background()

	tkn, err := r.resetPasswordTokenRepo.FindToken(ctx, token)
	if err != nil {
		return nil, err
	}

	user, errUser := r.userRepo.FindOneUser(ctx, tkn.UserID)
	if errUser != nil {
		return nil, errUser
	}

	if time.Now().After(tkn.ExpiresAt) {
		return nil, exception.NewHTTPError(400, "token is expired")
	}

	if payload.NewPassword != payload.ConfirmPassword {
		return nil, exception.
			NewHTTPError(400, "new and confirm password is not the same")
	}

	var errHash error
	payload.NewPassword, errHash = r.passwordHashService.
		HashPassword(payload.NewPassword)
	if errHash != nil {
		return nil, errHash
	}

	if err := r.userRepo.UpdatePassword(
		ctx,
		user.ID,
		payload.NewPassword,
	); err != nil {
		return nil, err
	}

	if err := r.resetPasswordTokenRepo.DeleteToken(
		ctx,
		token,
	); err != nil {
		return nil, err
	}

	response := dto.ResetPasswordResponse{
		Message: "your password has been reset",
	}

	return &response, nil
}
