package dto

type AuthDto struct {
	Username string `json:"username" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"required"`
}

type ChangePassword struct {
	Token           string `json:"token" binding:"required" validate:"required"`
	Password        string `json:"password" binding:"required" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required" validate:"required"`
}
