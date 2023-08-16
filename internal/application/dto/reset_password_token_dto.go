package dto

type GetResetPasswordTokenRequest struct {
	Email string `json:"email" valid:"email~invalid email"`
}

type GetResetPasswordTokenResponse struct {
	Message string `json:"message"`
}

type ResetPasswordRequest struct {
	NewPassword     string `json:"new_password"     valid:"required~new password is required,stringlength(8|128)~new password length should be between 8 - 128 chars"`
	ConfirmPassword string `json:"confirm_password" valid:"required~confim password is required,stringlength(8|128)~confirm password length should be between 8 - 128 chars"`
}

type ResetPasswordResponse GetResetPasswordTokenResponse
