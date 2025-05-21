package controller

import (
	"log"

	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/service"
)

type cronRecoveryCode struct {
	recoveryCodeService service.RecoveryService
}

func NewCronRecoveryCode() *cronRecoveryCode {
	return &cronRecoveryCode{
		recoveryCodeService: *service.NewRecoveryService(
			sqlRecoveryRepository,
			sqlUserRepository,
			*recoveryTokenService,
		),
	}
}

func (cronRecoveryCode *cronRecoveryCode) CheckRecoveryCodes() {
	err := cronRecoveryCode.recoveryCodeService.RecoveryCodeExpiry()
	if err != nil {
		log.Println("Failed expiry recovery codes", err)
	}

}
