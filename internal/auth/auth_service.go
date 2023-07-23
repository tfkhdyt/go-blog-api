package auth

import (
	"github.com/gofiber/fiber/v2"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/auth"
	"codeberg.org/tfkhdyt/blog-api/internal/domain/user"
)

type authService struct {
	authRepo auth.AuthRepository
	userRepo user.UserRepository
}

func NewAuthService(authRepo auth.AuthRepository, userRepo user.UserRepository) *authService {
	return &authService{authRepo, userRepo}
}

func (a *authService) Login(payload *auth.LoginRequest) (*auth.LoginResponse, error) {
	if err := payload.Validate(); err != nil {
		return nil, err
	}

	user, err := a.userRepo.FindOneUserByEmail(payload.Email)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	if err := user.ComparePassword(payload.Password); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	accessToken, errAccessToken := user.CreateAccessToken()
	if errAccessToken != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errAccessToken.Error())
	}

	refreshToken, errRefreshToken := user.CreateRefreshToken()
	if errRefreshToken != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errRefreshToken.Error())
	}

	if _, err := a.authRepo.AddToken(&auth.Auth{
		RefreshToken: *refreshToken,
	}); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := auth.LoginResponse{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}

	return &response, nil
}

func (a *authService) Refresh(
	userId uint,
	payload *auth.RefreshRequest,
) (*auth.RefreshResponse, error) {
	if err := payload.Validate(); err != nil {
		return nil, err
	}

	user, err := a.userRepo.FindOneUser(userId)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	if err := a.authRepo.VerifyToken(payload.RefreshToken); err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	accessToken, errAccessToken := user.CreateAccessToken()
	if errAccessToken != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errAccessToken.Error())
	}

	response := auth.RefreshResponse{
		AccessToken: *accessToken,
	}

	return &response, nil
}

func (a *authService) Logout(refreshToken string) (*auth.LogoutResponse, error) {
	if err := a.authRepo.RemoveToken(refreshToken); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := auth.LogoutResponse{
		Message: "refresh token has been deleted",
	}

	return &response, nil
}
