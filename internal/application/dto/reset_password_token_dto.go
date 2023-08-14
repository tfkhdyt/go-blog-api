package dto

type GetResetPasswordTokenRequest struct {
	Email string `json:"email" valid:"email~invalid email"`
}

type GetResetPasswordTokenResponse struct {
	Message string `json:"message"`
}
