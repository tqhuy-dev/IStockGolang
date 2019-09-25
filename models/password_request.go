package models

type ChangePasswordReq struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}