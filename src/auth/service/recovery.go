package service

import (
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/dto"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/repository/recovery_codes_repository"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/repository/user_repository"
	"github.com/CPU-commits/Template_Go-EventDriven/src/utils"
)

type RecoveryService struct {
	recoveryRepository   recovery_codes_repository.RecoveryRepository
	recoveryTokenService RecoveryTokenService
	userRepository       user_repository.UserRepository
}

func (recoveryService *RecoveryService) RecoveryCode(email string) (*model.Recovery, error) {
	randomString, err := utils.GenerateRandomString(6)
	if err != nil {
		return nil, err
	}
	user, err := recoveryService.userRepository.FindOneByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserLoginNotFound
	}

	recoveryCode := model.Recovery{
		Code:   randomString,
		IDUser: user.ID,
	}
	code, err := recoveryService.recoveryRepository.InsertOne(recoveryCode)
	if err != nil {
		return nil, err
	}

	return code, nil
}

func (recoveryService *RecoveryService) VerifyRecoveryCode(verifyRecoveryCode *dto.VerifyRecoveryCode) (bool, string, error) {

	user, err := recoveryService.userRepository.FindOneByEmail(verifyRecoveryCode.Email)
	if err != nil {
		return false, "", err
	}
	if user == nil {
		return false, "", ErrUserLoginNotFound
	}

	criteria := &recovery_codes_repository.RecoveryCriteria{
		Code:     verifyRecoveryCode.Code,
		IDUser:   user.ID,
		IsActive: utils.Bool(true),
	}

	codeExists, err := recoveryService.recoveryRepository.Exists(criteria)
	if err != nil {
		return false, "", err
	}
	if !codeExists {
		return false, "", ErrCodeNotValid
	}
	recoveryCode, err := recoveryService.recoveryRepository.FindOne(criteria)
	if err != nil {
		return false, "", err
	}
	err = utils.VerifyNotExpiredAt(recoveryCode.Expires_at, "utc", ErrCodeNotValid)
	if err != nil {
		return false, "", err
	}
	err = recoveryService.recoveryRepository.UpdateOne(recoveryCode.ID, recovery_codes_repository.RecoveryDataUpdate{
		IsActive: utils.Bool(false),
	})
	if err != nil {
		return false, "", err
	}

	token, err := recoveryService.recoveryTokenService.NewRecoveryToken(*user)
	if err != nil {
		return false, "", err
	}
	return codeExists, token, nil
}

func (recoveryService *RecoveryService) RecoveryCodeExpiry() error {
	recoveryCodes, err := recoveryService.recoveryRepository.Find(&recovery_codes_repository.RecoveryCriteria{
		IsActive: utils.Bool(true),
	})
	if err != nil {
		return err
	}

	return utils.ForEach(recoveryCodes, func(recovery model.Recovery) error {
		err := utils.VerifyNotExpiredAt(recovery.Expires_at, "utc", ErrCodeNotValid)
		if err == ErrCodeNotValid {
			errUpdateOne := recoveryService.recoveryRepository.UpdateOne(recovery.ID, recovery_codes_repository.RecoveryDataUpdate{
				IsActive: utils.Bool(false),
			})
			if errUpdateOne != nil {
				return errUpdateOne
			}
			return nil
		}
		if err != nil {
			return err
		}
		return nil
	})
}

func NewRecoveryService(
	recoveryRepository recovery_codes_repository.RecoveryRepository,
	userRepository user_repository.UserRepository,
	recoveryTokenService RecoveryTokenService,
) *RecoveryService {
	return &RecoveryService{
		recoveryRepository:   recoveryRepository,
		userRepository:       userRepository,
		recoveryTokenService: recoveryTokenService,
	}

}
