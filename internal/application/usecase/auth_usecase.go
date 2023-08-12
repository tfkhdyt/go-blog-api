package usecase

import (
	"codeberg.org/tfkhdyt/blog-api/internal/application/dto"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/repository"
	"codeberg.org/tfkhdyt/blog-api/internal/infrastructure/security"
)

type AuthUsecase struct {
	authRepo      repository.AuthRepository
	userRepo      repository.UserRepository
	bcryptService security.PasswordHashService
	jwtService    security.AuthTokenService
}

func (a *AuthUsecase) Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashedPassword, err := a.bcryptService.HashPassword(payload.Password)
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
		ID:        registeredUser.ID,
		FullName:  registeredUser.FullName,
		Username:  registeredUser.Username,
		Email:     registeredUser.Email,
		Role:      registeredUser.Role,
		CreatedAt: registeredUser.CreatedAt,
	}

	return &response, nil
}

func (a *AuthUsecase) Login(payload *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := a.userRepo.FindOneUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if err := a.bcryptService.ComparePassword(user.Password, payload.Password); err != nil {
		return nil, err
	}

	accessToken, errAccessToken := a.jwtService.CreateAccessToken(user.ID, user.Role)
	if errAccessToken != nil {
		return nil, errAccessToken
	}

	refreshToken, errRefreshToken := a.jwtService.CreateRefreshToken(user.ID, user.Role)
	if errRefreshToken != nil {
		return nil, errRefreshToken
	}

	if _, err := a.authRepo.AddToken(&entity.Auth{
		RefreshToken: refreshToken,
	}); err != nil {
		return nil, err
	}

	response := dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
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

	accessToken, errAccessToken := a.jwtService.CreateAccessToken(user.ID, user.Role)
	if errAccessToken != nil {
		return nil, errAccessToken
	}

	response := dto.RefreshResponse{
		AccessToken: accessToken,
	}

	return &response, nil
}

func (a *AuthUsecase) Logout(refreshToken string) (*dto.LogoutResponse, error) {
	if err := a.authRepo.RemoveToken(refreshToken); err != nil {
		return nil, err
	}

	response := dto.LogoutResponse{
		Message: "refresh token has been deleted",
	}

	return &response, nil
}
