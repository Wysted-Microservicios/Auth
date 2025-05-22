package service

import (
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/dto"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/repository/auth_repository"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/repository/recovery_tokens_repository"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/repository/user_repository"
	"github.com/CPU-commits/Template_Go-EventDriven/src/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepository          auth_repository.AuthRepository
	userRepository          user_repository.UserRepository
	recoveryTokenService    RecoveryTokenService
	recoteryTokenRepository recovery_tokens_repository.RecoveryTokenRepository
}

func (authService *AuthService) Register(
	registerDto *dto.RegisterDto,
) error {
	user, err := registerDto.ToModel()
	if err != nil {
		return err
	}
	existsEmailOrUsername, err := authService.userRepository.Exists(&user_repository.Criteria{
		Or: []user_repository.Criteria{
			{
				Email: registerDto.Email,
			},
			{
				Username: registerDto.Username,
			},
		},
	})
	if existsEmailOrUsername {
		return ErrExistsEmailOrUsername
	}
	if err != nil {
		return err
	}
	_, err = authService.userRepository.InsertOne(user, registerDto.Password)
	return err
}

func (authService *AuthService) Login(authDto dto.AuthDto) (*model.User, int64, error) {
	auth, err := authService.authRepository.FindOneByUsername(authDto.Username)
	if err != nil {
		return nil, 0, utils.ErrRepositoryFailed
	}
	if auth == nil {
		return nil, 0, ErrUserLoginNotFound
	}
	if err := bcrypt.CompareHashAndPassword(
		[]byte(auth.Password),
		[]byte(authDto.Password),
	); err != nil {
		return nil, 0, ErrInvalidCredentials
	}
	// User
	user, err := authService.userRepository.FindOneByEmail(authDto.Username)
	if err != nil {
		return nil, 0, err
	}
	if user == nil {
		return nil, 0, ErrUserLoginNotFound
	}

	return user, auth.ID, nil
}

func (authService *AuthService) ChangePassword(changePasswordDto dto.ChangePassword) error {
	if changePasswordDto.Password != changePasswordDto.ConfirmPassword {
		return ErrInvalidCredentials
	}

	tokenData, err := authService.recoveryTokenService.CheckToken(changePasswordDto.Token)
	if err != nil {
		return err
	}
	err = authService.authRepository.UpdateOne(tokenData.IDUser, &auth_repository.AuthDataUpdate{
		Password: changePasswordDto.Password,
	})
	if err != nil {
		return err
	}
	err = authService.recoteryTokenRepository.UpdateOne(tokenData.ID, recovery_tokens_repository.RecoveryTokenUpdate{
		IsUsed: utils.Bool(true),
	})
	if err != nil {
		return err
	}
	return nil
}

func NewAuthService(
	authRepository auth_repository.AuthRepository,
	userRepository user_repository.UserRepository,
	recoveryTokenService RecoveryTokenService,
	recoteryTokenRepository recovery_tokens_repository.RecoveryTokenRepository,

) *AuthService {
	return &AuthService{
		authRepository:          authRepository,
		userRepository:          userRepository,
		recoveryTokenService:    recoveryTokenService,
		recoteryTokenRepository: recoteryTokenRepository,
	}
}
