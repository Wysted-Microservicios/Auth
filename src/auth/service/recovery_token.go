package service

import (
	"time"

	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/repository/recovery_tokens_repository"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/repository/token_generator_repository"
	"github.com/CPU-commits/Template_Go-EventDriven/src/utils"
)

const (
	maxRecoveryTokenTime = time.Minute * 5
)

type RecoveryTokenService struct {
	recoveryTokenRepository recovery_tokens_repository.RecoveryTokenRepository
	recoveryTokenGenerator  token_generator_repository.TokenGenerator
}

func (recoveryTokenService *RecoveryTokenService) NewRecoveryToken(user model.User) (string, error) {
	expiredAt := time.Now().UTC().Add(maxRecoveryTokenTime)

	token, err := recoveryTokenService.recoveryTokenGenerator.NewRecoveryCodeToken(
		expiredAt,
		user,
	)
	if err != nil {
		return "", err
	}

	_, err = recoveryTokenService.recoveryTokenRepository.InsertOne(model.RecoveryToken{
		Token:      token,
		IDUser:     user.ID,
		Expires_at: expiredAt,
	})
	if err != nil {
		return "", err
	}
	return token, nil
}
func (recoveryTokenService *RecoveryTokenService) CheckToken(token string) (model.RecoveryToken, error) {
	criteria := &recovery_tokens_repository.RecoveryTokenCriteria{
		Token:  token,
		IsUsed: utils.Bool(false),
	}
	check, err := recoveryTokenService.recoveryTokenRepository.Exists(criteria)
	if err != nil {
		return model.RecoveryToken{}, err
	}
	if !check {
		return model.RecoveryToken{}, ErrTokenNotValid
	}
	checkToken, err := recoveryTokenService.recoveryTokenRepository.FindOne(criteria)
	if err != nil {
		return model.RecoveryToken{}, err
	}
	err = utils.VerifyNotExpiredAt(checkToken.Expires_at, "utc", ErrTokenNotValid)
	if err != nil {
		return model.RecoveryToken{}, err
	}

	return *checkToken, nil
}

func (recoveryTokenService *RecoveryTokenService) RecoveryTokenExpiry() error {
	recoveryTokens, err := recoveryTokenService.recoveryTokenRepository.Find(&recovery_tokens_repository.RecoveryTokenCriteria{
		IsUsed: utils.Bool(false),
	})
	if err != nil {
		return err
	}
	return utils.ForEach(recoveryTokens, func(recoveryToken model.RecoveryToken) error {
		err := utils.VerifyNotExpiredAt(recoveryToken.Expires_at, "utc", ErrTokenNotValid)
		if err == ErrTokenNotValid {
			errUpdateOne := recoveryTokenService.recoveryTokenRepository.UpdateOne(recoveryToken.ID, recovery_tokens_repository.RecoveryTokenUpdate{
				IsUsed: utils.Bool(true),
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

func NewRecoveryTokenService(
	recoveryTokenRepository recovery_tokens_repository.RecoveryTokenRepository,
	recoveryTokenGenerator token_generator_repository.TokenGenerator,

) *RecoveryTokenService {
	return &RecoveryTokenService{
		recoveryTokenRepository: recoveryTokenRepository,
		recoveryTokenGenerator:  recoveryTokenGenerator,
	}

}
