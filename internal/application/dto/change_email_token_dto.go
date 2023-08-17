package dto

type GetChangeEmailTokenRequest struct {
	NewEmail string `json:"new_email" valid:"email~invalid new email"`
	Password string `json:"password"  valid:"required~new password is required,stringlength(8|128)~new password length should be between 8 - 128 chars"`
}

type GetChangeEmailTokenResponse struct {
	Message string `json:"message"`
}

type ChangeEmailResponse GetChangeEmailTokenResponse
