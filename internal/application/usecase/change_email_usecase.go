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

type ChangeEmailUsecase struct {
	changeEmailTokenRepo repository.ChangeEmailTokenRepository `di.inject:"changeEmailRequestRepo"`
	userRepo             repository.UserRepository             `di.inject:"userRepo"`
	idService            service.IDService                     `di.inject:"idService"`
	emailService         service.EmailService                  `di.inject:"emailService"`
	passwordHashService  service.PasswordHashService           `di.inject:"passwordHashService"`
}

func (r *ChangeEmailUsecase) GetChangeEmailToken(
	userId int32,
	payload *dto.GetChangeEmailTokenRequest,
) (*dto.GetChangeEmailTokenResponse, error) {
	ctx := context.Background()

	user, err := r.userRepo.FindOneUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	if err := r.passwordHashService.ComparePassword(
		user.Password,
		payload.Password,
	); err != nil {
		return nil, err
	}

	if _, err := r.userRepo.FindOneUserByEmail(
		ctx,
		payload.NewEmail,
	); err == nil {
		return nil, exception.
			NewHTTPError(400, "this email has been used by other user")
	}

	token := r.idService.GenerateID()

	if _, errAddToken := r.changeEmailTokenRepo.AddToken(
		ctx,
		userId,
		&entity.ChangeEmailRequest{
			NewEmail:  payload.NewEmail,
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
			Email: payload.NewEmail,
			Name:  user.FullName,
		},
		"Change-Email Token",
		fmt.Sprintf("Your change-email token is <b>%v<b>", token),
	); err != nil {
		return nil, err
	}

	response := dto.GetChangeEmailTokenResponse{
		Message: "change-email token has been sent to your new email.",
	}

	return &response, nil
}

func (r *ChangeEmailUsecase) ChangeEmail(
	token string,
) (*dto.ChangeEmailResponse, error) {
	ctx := context.Background()

	request, err := r.changeEmailTokenRepo.FindToken(ctx, token)
	if err != nil {
		return nil, err
	}

	user, errUser := r.userRepo.FindOneUser(ctx, request.UserID)
	if errUser != nil {
		return nil, errUser
	}

	if time.Now().After(request.ExpiresAt) {
		return nil, exception.NewHTTPError(400, "token is expired")
	}

	if _, err := r.userRepo.UpdateEmail(
		ctx,
		user.ID,
		request.NewEmail,
	); err != nil {
		return nil, err
	}

	if err := r.changeEmailTokenRepo.DeleteToken(
		ctx,
		token,
	); err != nil {
		return nil, err
	}

	response := dto.ChangeEmailResponse{
		Message: "your email has been changed",
	}

	return &response, nil
}
