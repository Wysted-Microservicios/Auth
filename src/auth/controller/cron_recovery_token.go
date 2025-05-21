package controller

import (
	"log"

	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/service"
)

type CronRecoveryToken struct {
	recoveryTokenService service.RecoveryTokenService
}

func NewCronRecoveryToken() *CronRecoveryToken {
	return &CronRecoveryToken{
		recoveryTokenService: *service.NewRecoveryTokenService(
			sqlRecoveryTokenRepository,
			generator,
		),
	}
}

func (cronRecoveryToken *CronRecoveryToken) CheckRecoveryTokens() {
	err := cronRecoveryToken.recoveryTokenService.RecoveryTokenExpiry()
	if err != nil {
		log.Println("Failed expiry recovery codes", err)
	}

}
