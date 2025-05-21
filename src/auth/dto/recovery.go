package dto

type RecoveryDto struct {
	Code     string `json:"code" binding:"required" validate:"required"`
	IDUser   int64  `json:"idUser" binding:"required" validate:"required"`
	IsActive bool   `json:"isActive" binding:"required" validate:"required"`
}

type GetRecoveryCode struct {
	Email string `json:"email" binding:"required" validate:"required"`
}

type VerifyRecoveryCode struct {
	Email string `json:"email" binding:"required" validate:"required"`
	Code  string `json:"code" binding:"required" validate:"required"`
}
