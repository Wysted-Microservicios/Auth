package auth

import (
	"os"
	"testing"

	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/dto"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/service"
	"github.com/CPU-commits/Template_Go-EventDriven/src/utils"
	"github.com/CPU-commits/Template_Go-EventDriven/tests"
	"github.com/CPU-commits/Template_Go-EventDriven/tests/auth/auth/mock_struct"
	"github.com/CPU-commits/Template_Go-EventDriven/tests/mocks"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

var AuthS *service.AuthService
var RecoveryTokenS *service.RecoveryTokenService

var mockUserRepo *mocks.MockUserRepository
var mockTokenGeneratorRepo *mocks.MockTokenGenerator
var mockRecoveryTokenRepo *mocks.MockRecoveryTokenRepository
var mockAuthRepo *mocks.MockAuthRepository

func TestMain(m *testing.M) {
	mockUserRepo = &mocks.MockUserRepository{}
	mockTokenGeneratorRepo = &mocks.MockTokenGenerator{}
	mockRecoveryTokenRepo = &mocks.MockRecoveryTokenRepository{}
	mockAuthRepo = &mocks.MockAuthRepository{}

	RecoveryTokenS = service.NewRecoveryTokenService(mockRecoveryTokenRepo, mockTokenGeneratorRepo, mockUserRepo)

	AuthS = service.NewAuthService(mockAuthRepo, mockUserRepo, *RecoveryTokenS, mockRecoveryTokenRepo)

	code := m.Run()
	os.Exit(code)
}

func TestRegister(t *testing.T) {
	testCases := []struct {
		Name        string
		RegisterDto *dto.RegisterDto
		Repository  struct {
			Exists    bool
			InsertOne *model.User
		}
		ExpectedReturns struct {
			RepoError_1   error
			RepoError_2   error
			ExpectedError error
		}
	}{
		{
			Name:        "Register Ok",
			RegisterDto: tests.RegisterDto_1,
			Repository: mock_struct.RegisterRepository{
				Exists:    false,
				InsertOne: tests.User_1,
			},
			ExpectedReturns: mock_struct.RegisterExpectedReturns{
				RepoError_1:   nil,
				RepoError_2:   nil,
				ExpectedError: nil,
			},
		},
		{
			Name:        "Register Email/Username exists",
			RegisterDto: tests.RegisterDto_1,
			Repository: mock_struct.RegisterRepository{
				Exists:    true,
				InsertOne: nil,
			},
			ExpectedReturns: mock_struct.RegisterExpectedReturns{
				RepoError_1:   nil,
				RepoError_2:   nil,
				ExpectedError: service.ErrExistsEmailOrUsername,
			},
		},
		{
			Name:        "Exists RepositoryFailed",
			RegisterDto: tests.RegisterDto_1,
			Repository: mock_struct.RegisterRepository{
				Exists:    false,
				InsertOne: nil,
			},
			ExpectedReturns: mock_struct.RegisterExpectedReturns{
				RepoError_1:   utils.ErrRepositoryFailed,
				RepoError_2:   nil,
				ExpectedError: utils.ErrRepositoryFailed,
			},
		},
		{
			Name:        "Exists RepositoryFailed",
			RegisterDto: tests.RegisterDto_1,
			Repository: mock_struct.RegisterRepository{
				Exists:    false,
				InsertOne: nil,
			},
			ExpectedReturns: mock_struct.RegisterExpectedReturns{
				RepoError_1:   nil,
				RepoError_2:   utils.ErrRepositoryFailed,
				ExpectedError: utils.ErrRepositoryFailed,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mockUserRepo.ExpectedCalls = nil

			mockUserRepo.On(tests.EXISTS,
				mock.AnythingOfType(tests.USER_CRITERIA_PTR),
			).Return(tc.Repository.Exists, tc.ExpectedReturns.RepoError_1)

			mockUserRepo.On(tests.INSERT_ONE,
				mock.AnythingOfType(tests.USER_MODEL_PTR),
				mock.AnythingOfType(tests.STRING),
			).Return(tc.Repository.InsertOne, tc.ExpectedReturns.RepoError_2)

			err := AuthS.Register(tc.RegisterDto)
			mockUserRepo.AssertExpectations(t)

			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)
		})
	}
}

func TestLogin(t *testing.T) {
	testCases := []struct {
		Name       string
		AuthDto    dto.AuthDto
		Repository struct {
			FindOneByUsername *model.Auth
			FindOneByEmail    *model.User
		}
		ExpectedReturns struct {
			RepoError_1      error
			RepoError_2      error
			ExpectedReturn_1 *model.User
			ExpectedReturn_2 int64
			ExpectedError    error
		}
	}{
		{
			Name:    "Log in success",
			AuthDto: *tests.AuthDto_1,
			Repository: mock_struct.LoginRepository{
				FindOneByUsername: tests.Auth_1,
				FindOneByEmail:    tests.User_1,
			},
			ExpectedReturns: mock_struct.LoginExpectedReturns{
				RepoError_1:      nil,
				RepoError_2:      nil,
				ExpectedReturn_1: tests.User_1,
				ExpectedReturn_2: tests.Auth_1.ID,
				ExpectedError:    nil,
			},
		},
		{
			Name:    "User not found to log in",
			AuthDto: *tests.AuthDto_1,
			Repository: mock_struct.LoginRepository{
				FindOneByUsername: nil,
				FindOneByEmail:    nil,
			},
			ExpectedReturns: mock_struct.LoginExpectedReturns{
				RepoError_1:      nil,
				RepoError_2:      nil,
				ExpectedReturn_1: nil,
				ExpectedReturn_2: 0,
				ExpectedError:    service.ErrUserLoginNotFound,
			},
		},
		{
			Name:    "Invalid Credentials",
			AuthDto: *tests.AuthDto_1,
			Repository: mock_struct.LoginRepository{
				FindOneByUsername: tests.Auth_2,
				FindOneByEmail:    nil,
			},
			ExpectedReturns: mock_struct.LoginExpectedReturns{
				RepoError_1:      nil,
				RepoError_2:      nil,
				ExpectedReturn_1: nil,
				ExpectedReturn_2: 0,
				ExpectedError:    service.ErrInvalidCredentials,
			},
		},
		{
			Name:    "User not found",
			AuthDto: *tests.AuthDto_1,
			Repository: mock_struct.LoginRepository{
				FindOneByUsername: tests.Auth_1,
				FindOneByEmail:    nil,
			},
			ExpectedReturns: mock_struct.LoginExpectedReturns{
				RepoError_1:      nil,
				RepoError_2:      nil,
				ExpectedReturn_1: nil,
				ExpectedReturn_2: 0,
				ExpectedError:    service.ErrUserLoginNotFound,
			},
		},
		{
			Name:    "FindOneByUsername RepositoryFailed",
			AuthDto: *tests.AuthDto_1,
			Repository: mock_struct.LoginRepository{
				FindOneByUsername: nil,
				FindOneByEmail:    nil,
			},
			ExpectedReturns: mock_struct.LoginExpectedReturns{
				RepoError_1:      utils.ErrRepositoryFailed,
				RepoError_2:      nil,
				ExpectedReturn_1: nil,
				ExpectedReturn_2: 0,
				ExpectedError:    utils.ErrRepositoryFailed,
			},
		},
		{
			Name:    "FindOneByEmail RepositoryFailed",
			AuthDto: *tests.AuthDto_1,
			Repository: mock_struct.LoginRepository{
				FindOneByUsername: tests.Auth_1,
				FindOneByEmail:    nil,
			},
			ExpectedReturns: mock_struct.LoginExpectedReturns{
				RepoError_1:      nil,
				RepoError_2:      utils.ErrRepositoryFailed,
				ExpectedReturn_1: nil,
				ExpectedReturn_2: 0,
				ExpectedError:    utils.ErrRepositoryFailed,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mockUserRepo.ExpectedCalls = nil
			mockAuthRepo.ExpectedCalls = nil

			mockAuthRepo.On(tests.FIND_ONE_BY_USERNAME,
				mock.AnythingOfType(tests.STRING),
			).Return(tc.Repository.FindOneByUsername, tc.ExpectedReturns.RepoError_1)

			mockUserRepo.On(tests.FIND_ONE_BY_EMAIL,
				mock.AnythingOfType(tests.STRING),
			).Maybe().Return(tc.Repository.FindOneByEmail, tc.ExpectedReturns.RepoError_2)

			result_1, result_2, err := AuthS.Login(tc.AuthDto)

			mockUserRepo.AssertExpectations(t)
			mockAuthRepo.AssertExpectations(t)

			assert.Equal(t, tc.ExpectedReturns.ExpectedReturn_1, result_1)
			assert.Equal(t, tc.ExpectedReturns.ExpectedReturn_2, result_2)
			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)

		})
	}
}

func TestChangePassword(t *testing.T) {
	testCases := []struct {
		Name              string
		ChangePasswordDto dto.ChangePassword

		Repository struct {
			RecoveryTokenExists    bool
			RecoveryTokenFindOne   *model.RecoveryToken
			AuthUpdateOne          error
			RecoveryTokenUpdateOne error
		}
		ExpectedReturns struct {
			RepoError_1   error
			RepoError_2   error
			RepoError_3   error
			RepoError_4   error
			ExpectedError error
		}
	}{
		{
			Name:              "Invalid Credentials",
			ChangePasswordDto: tests.ChangePasswordDto_2,
			Repository:        mock_struct.ChangePasswordRepository{},
			ExpectedReturns: mock_struct.ChangePasswordExpectedReturns{
				RepoError_1:   nil,
				RepoError_2:   nil,
				RepoError_3:   nil,
				RepoError_4:   nil,
				ExpectedError: service.ErrInvalidCredentials,
			},
		},
		{
			Name:              "Change password success",
			ChangePasswordDto: tests.ChangePasswordDto_1,
			Repository: mock_struct.ChangePasswordRepository{
				RecoveryTokenExists:    true,
				RecoveryTokenFindOne:   tests.RecoveryToken_1,
				AuthUpdateOne:          nil,
				RecoveryTokenUpdateOne: nil,
			},
			ExpectedReturns: mock_struct.ChangePasswordExpectedReturns{
				RepoError_1:   nil,
				RepoError_2:   nil,
				RepoError_3:   nil,
				RepoError_4:   nil,
				ExpectedError: nil,
			},
		},
		{
			Name:              "Invalid Token",
			ChangePasswordDto: tests.ChangePasswordDto_1,
			Repository: mock_struct.ChangePasswordRepository{
				RecoveryTokenExists:    false,
				RecoveryTokenFindOne:   nil,
				AuthUpdateOne:          nil,
				RecoveryTokenUpdateOne: nil,
			},
			ExpectedReturns: mock_struct.ChangePasswordExpectedReturns{
				RepoError_1:   nil,
				RepoError_2:   nil,
				RepoError_3:   nil,
				RepoError_4:   nil,
				ExpectedError: service.ErrTokenNotValid,
			},
		},
		{
			Name:              "Invalid Token (Token Expired)",
			ChangePasswordDto: tests.ChangePasswordDto_1,
			Repository: mock_struct.ChangePasswordRepository{
				RecoveryTokenExists:    true,
				RecoveryTokenFindOne:   tests.RecoveryToken_2,
				AuthUpdateOne:          nil,
				RecoveryTokenUpdateOne: nil,
			},
			ExpectedReturns: mock_struct.ChangePasswordExpectedReturns{
				RepoError_1:   nil,
				RepoError_2:   nil,
				RepoError_3:   nil,
				RepoError_4:   nil,
				ExpectedError: service.ErrTokenNotValid,
			},
		},
		{
			Name:              "Exists RepositoryFailed",
			ChangePasswordDto: tests.ChangePasswordDto_1,
			Repository: mock_struct.ChangePasswordRepository{
				RecoveryTokenExists:    false,
				RecoveryTokenFindOne:   nil,
				AuthUpdateOne:          nil,
				RecoveryTokenUpdateOne: nil,
			},
			ExpectedReturns: mock_struct.ChangePasswordExpectedReturns{
				RepoError_1:   utils.ErrRepositoryFailed,
				RepoError_2:   nil,
				ExpectedError: utils.ErrRepositoryFailed,
			},
		},
		{
			Name:              "FindOne RepositoryFailed",
			ChangePasswordDto: tests.ChangePasswordDto_1,
			Repository: mock_struct.ChangePasswordRepository{
				RecoveryTokenExists:    true,
				RecoveryTokenFindOne:   nil,
				AuthUpdateOne:          nil,
				RecoveryTokenUpdateOne: nil,
			},
			ExpectedReturns: mock_struct.ChangePasswordExpectedReturns{
				RepoError_1:   nil,
				RepoError_2:   utils.ErrRepositoryFailed,
				ExpectedError: utils.ErrRepositoryFailed,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mockAuthRepo.ExpectedCalls = nil
			mockRecoveryTokenRepo.ExpectedCalls = nil

			mockRecoveryTokenRepo.On(tests.EXISTS,
				mock.AnythingOfType(tests.RECOVERY_TOKEN_CRITERIA_PTR),
			).Return(tc.Repository.RecoveryTokenExists, tc.ExpectedReturns.RepoError_1).Maybe()

			mockRecoveryTokenRepo.On(tests.FIND_ONE,
				mock.AnythingOfType(tests.RECOVERY_TOKEN_CRITERIA_PTR),
			).Return(tc.Repository.RecoveryTokenFindOne, tc.ExpectedReturns.RepoError_2).Maybe()

			mockAuthRepo.On(tests.UPDATE_ONE,
				mock.AnythingOfType(tests.INT64),
				mock.AnythingOfType(tests.AUTH_DATA_UPDATE),
			).Return(tc.Repository.AuthUpdateOne).Maybe()

			mockRecoveryTokenRepo.On(tests.UPDATE_ONE,
				mock.AnythingOfType(tests.INT64),
				mock.AnythingOfType(tests.RECOVERY_TOKEN_DATA_UPDATE),
			).Return(tc.Repository.RecoveryTokenUpdateOne).Maybe()

			err := AuthS.ChangePassword(tc.ChangePasswordDto)

			mockAuthRepo.AssertExpectations(t)
			mockRecoveryTokenRepo.AssertExpectations(t)

			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)
		})
	}
}

// func Test{nombre de la funcion a testear}(t *testing.T) {
// 	testCases := []struct {
// 		Name string
// 	}{}
// 	for _, tc := range testCases {
// 		t.Run(tc.Name, func(t *testing.T) {
// 			Resetear expectativas de los mocks antes de cada test
// 			mockUserRepo.ExpectedCalls = nil

// 			Probar funcion
// 			err := s.UserExist(tc.UserID)
// 			mockUserRepo.AssertExpectations(t)

// 			Compara el error entregado, por el error esperado
// 			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)

// 		})
// 	}
// }
