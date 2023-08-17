package usecase

import (
	"codeberg.org/tfkhdyt/blog-api/internal/application/dto"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/repository"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/service"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type AuthUsecase struct {
	authRepo            repository.AuthRepository   `di.inject:"authRepo"`
	userRepo            repository.UserRepository   `di.inject:"userRepo"`
	passwordHashService service.PasswordHashService `di.inject:"passwordHashService"`
	authTokenService    service.AuthTokenService    `di.inject:"authTokenService"`
}

func (a *AuthUsecase) Register(
	payload *dto.RegisterRequest,
) (*dto.RegisterResponse, error) {
	hashedPassword, err := a.passwordHashService.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	registeredUser, errRegister := a.userRepo.Register(&entity.User{
		FullName: payload.FullName,
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashedPassword,
	})
	if errRegister != nil {
		return nil, errRegister
	}

	response := dto.RegisterResponse{
		Message: "your account registration has been successful",
		Data: dto.RegisterResponseData{
			ID:        registeredUser.ID,
			FullName:  registeredUser.FullName,
			Username:  registeredUser.Username,
			Email:     registeredUser.Email,
			Role:      registeredUser.Role,
			CreatedAt: registeredUser.CreatedAt,
		},
	}

	return &response, nil
}

func (a *AuthUsecase) Login(
	payload *dto.LoginRequest,
) (*dto.LoginResponse, error) {
	user, err := a.userRepo.FindOneUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if err := a.passwordHashService.ComparePassword(
		user.Password,
		payload.Password,
	); err != nil {
		return nil, err
	}

	accessToken, errAccessToken := a.authTokenService.CreateAccessToken(
		user.ID,
		user.Role,
	)
	if errAccessToken != nil {
		return nil, errAccessToken
	}

	refreshToken, errRefreshToken := a.authTokenService.CreateRefreshToken(
		user.ID,
		user.Role,
	)
	if errRefreshToken != nil {
		return nil, errRefreshToken
	}

	if _, err := a.authRepo.AddToken(&entity.Auth{
		RefreshToken: refreshToken,
	}); err != nil {
		return nil, err
	}

	response := dto.LoginResponse{
		Message: "you have successfully logged in",
		Data: dto.LoginResponseData{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}

	return &response, nil
}

func (a *AuthUsecase) Refresh(
	userId uint,
	payload *dto.RefreshRequest,
) (*dto.RefreshResponse, error) {
	user, err := a.userRepo.FindOneUser(userId)
	if err != nil {
		return nil, err
	}

	if err := a.authRepo.VerifyToken(payload.RefreshToken); err != nil {
		return nil, err
	}

	accessToken, errAccessToken := a.authTokenService.CreateAccessToken(
		user.ID,
		user.Role,
	)
	if errAccessToken != nil {
		return nil, errAccessToken
	}

	response := dto.RefreshResponse{
		Message: "access token has been refreshed",
		Data: dto.RefreshResponseData{
			AccessToken: accessToken,
		},
	}

	return &response, nil
}

func (a *AuthUsecase) Logout(
	refreshToken string,
) (*dto.LogoutResponse, error) {
	if err := a.authRepo.RemoveToken(refreshToken); err != nil {
		return nil, err
	}

	response := dto.LogoutResponse{
		Message: "you've logged out successfully",
	}

	return &response, nil
}

func (a *AuthUsecase) ChangePassword(
	userId uint,
	payload *dto.ChangePasswordRequest,
) (*dto.ChangePasswordResponse, error) {
	user, err := a.userRepo.FindOneUser(userId)
	if err != nil {
		return nil, err
	}

	if err := a.passwordHashService.ComparePassword(
		user.Password,
		payload.OldPassword,
	); err != nil {
		return nil, err
	}

	if payload.NewPassword != payload.ConfirmPassword {
		return nil, exception.
			NewHTTPError(400, "new and confirm password is not the same")
	}

	hashedPassword, errHash := a.passwordHashService.HashPassword(
		payload.NewPassword,
	)
	if errHash != nil {
		return nil, errHash
	}

	if _, err := a.userRepo.UpdateUser(user, &entity.User{
		Password: hashedPassword,
	}); err != nil {
		return nil, err
	}

	response := dto.ChangePasswordResponse{
		Message: "your password has been changed",
	}

	return &response, nil
}
