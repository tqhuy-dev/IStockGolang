package models

type ChangePasswordReq struct {
	Email string `json:"email"`
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}