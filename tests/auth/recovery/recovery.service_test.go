package auth

import (
	"os"
	"testing"

	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/dto"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/service"
	"github.com/CPU-commits/Template_Go-EventDriven/src/utils"
	"github.com/CPU-commits/Template_Go-EventDriven/tests"
	"github.com/CPU-commits/Template_Go-EventDriven/tests/auth/recovery/mock_struct"
	"github.com/CPU-commits/Template_Go-EventDriven/tests/mocks"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

var RecoveryS *service.RecoveryService
var TokenS *service.RecoveryTokenService

var mockUserRepo *mocks.MockUserRepository
var mockTokenGeneratorRepo *mocks.MockTokenGenerator
var mockRecoveryTokenRepo *mocks.MockRecoveryTokenRepository
var mockRecoveryRepo *mocks.MockRecoveryRepository

func TestMain(m *testing.M) {
	mockUserRepo = &mocks.MockUserRepository{}
	mockTokenGeneratorRepo = &mocks.MockTokenGenerator{}
	mockRecoveryTokenRepo = &mocks.MockRecoveryTokenRepository{}
	mockRecoveryRepo = &mocks.MockRecoveryRepository{}

	TokenS = service.NewRecoveryTokenService(mockRecoveryTokenRepo, mockTokenGeneratorRepo)

	RecoveryS = service.NewRecoveryService(mockRecoveryRepo, mockUserRepo, *TokenS)

	code := m.Run()
	os.Exit(code)
}

func TestRecoveryCode(t *testing.T) {
	testCases := []struct {
		Name  string
		Email string

		Repository struct {
			FindOneByEmail    *model.User
			RecoveryInsertOne *model.Recovery
		}
		ExpectedReturns struct {
			RepoError_1    error
			RepoError_2    error
			ExpectedReturn *model.Recovery
			ExpectedError  error
		}
	}{
		{
			Name:  "Recovery code success",
			Email: tests.Email_1,
			Repository: mock_struct.RecoveryCodeRepository{
				FindOneByEmail:    tests.User_1,
				RecoveryInsertOne: tests.RecoveryCode_1,
			},
			ExpectedReturns: mock_struct.RecoveryCodeExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    nil,
				ExpectedReturn: tests.RecoveryCode_1,
				ExpectedError:  nil,
			},
		},
		{
			Name:  "Not User",
			Email: tests.Email_1,
			Repository: mock_struct.RecoveryCodeRepository{
				FindOneByEmail:    nil,
				RecoveryInsertOne: nil,
			},
			ExpectedReturns: mock_struct.RecoveryCodeExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    nil,
				ExpectedReturn: nil,
				ExpectedError:  service.ErrUserLoginNotFound,
			},
		},
		{
			Name:  "FindOneByEmail RepositoryFailed",
			Email: tests.Email_1,
			Repository: mock_struct.RecoveryCodeRepository{
				FindOneByEmail:    nil,
				RecoveryInsertOne: nil,
			},
			ExpectedReturns: mock_struct.RecoveryCodeExpectedReturns{
				RepoError_1:    utils.ErrRepositoryFailed,
				RepoError_2:    nil,
				ExpectedReturn: nil,
				ExpectedError:  utils.ErrRepositoryFailed,
			},
		},
		{
			Name:  "InsertOne RepositoryFailed",
			Email: tests.Email_1,
			Repository: mock_struct.RecoveryCodeRepository{
				FindOneByEmail:    tests.User_1,
				RecoveryInsertOne: nil,
			},
			ExpectedReturns: mock_struct.RecoveryCodeExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    utils.ErrRepositoryFailed,
				ExpectedReturn: nil,
				ExpectedError:  utils.ErrRepositoryFailed,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mockUserRepo.ExpectedCalls = nil
			mockRecoveryRepo.ExpectedCalls = nil

			mockUserRepo.On(tests.FIND_ONE_BY_EMAIL,
				mock.AnythingOfType(tests.STRING),
			).Return(tc.Repository.FindOneByEmail, tc.ExpectedReturns.RepoError_1)

			mockRecoveryRepo.On(tests.INSERT_ONE,
				mock.AnythingOfType(tests.RECOVERY_MODEL),
			).Return(tc.Repository.RecoveryInsertOne, tc.ExpectedReturns.RepoError_2)

			result, err := RecoveryS.RecoveryCode(tc.Email)

			mockUserRepo.AssertExpectations(t)
			mockRecoveryRepo.AssertExpectations(t)

			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)
			assert.Equal(t, tc.ExpectedReturns.ExpectedReturn, result)

		})
	}
}

func TestVerifyRecoveryCode(t *testing.T) {
	testCases := []struct {
		Name               string
		VerifyRecoveryCode *dto.VerifyRecoveryCode

		Repository struct {
			FindOneByEmail         *model.User
			RecoveryExists         bool
			RecoveryFindOne        *model.Recovery
			RecoveryTokenGenerator string
			RecoveryTokenInsertOne *model.RecoveryToken
		}
		ExpectedReturns struct {
			RepoError_1      error
			RepoError_2      error
			RepoError_3      error
			RepoError_4      error
			RepoError_5      error
			RepoError_6      error
			ExpectedReturn_1 bool
			ExpectedReturn_2 string
			ExpectedError    error
		}
	}{
		{
			Name:               "Recovery code is valid",
			VerifyRecoveryCode: tests.VerifyRecoveryCode_1,
			Repository: mock_struct.VerifyRecoveryCodeRepository{
				FindOneByEmail:         tests.User_1,
				RecoveryExists:         true,
				RecoveryFindOne:        tests.RecoveryCode_1,
				RecoveryTokenGenerator: tests.Token_1,
				RecoveryTokenInsertOne: tests.RecoveryToken_1,
			},
			ExpectedReturns: mock_struct.VerifyRecoveryCodeExpectedReturns{
				RepoError_1:      nil,
				RepoError_2:      nil,
				RepoError_3:      nil,
				RepoError_4:      nil,
				RepoError_5:      nil,
				RepoError_6:      nil,
				ExpectedReturn_1: true,
				ExpectedReturn_2: tests.Token_1,
				ExpectedError:    nil,
			},
		},
		{
			Name:               "User not found",
			VerifyRecoveryCode: tests.VerifyRecoveryCode_1,
			Repository: mock_struct.VerifyRecoveryCodeRepository{
				FindOneByEmail:         nil,
				RecoveryExists:         false,
				RecoveryFindOne:        nil,
				RecoveryTokenGenerator: "",
				RecoveryTokenInsertOne: nil,
			},
			ExpectedReturns: mock_struct.VerifyRecoveryCodeExpectedReturns{
				RepoError_1:      nil,
				RepoError_2:      nil,
				RepoError_3:      nil,
				RepoError_4:      nil,
				RepoError_5:      nil,
				RepoError_6:      nil,
				ExpectedReturn_1: false,
				ExpectedReturn_2: "",
				ExpectedError:    service.ErrUserLoginNotFound,
			},
		},
		{
			Name:               "Code not valid",
			VerifyRecoveryCode: tests.VerifyRecoveryCode_1,
			Repository: mock_struct.VerifyRecoveryCodeRepository{
				FindOneByEmail:         tests.User_1,
				RecoveryExists:         false,
				RecoveryFindOne:        nil,
				RecoveryTokenGenerator: "",
				RecoveryTokenInsertOne: nil,
			},
			ExpectedReturns: mock_struct.VerifyRecoveryCodeExpectedReturns{
				RepoError_1:      nil,
				RepoError_2:      nil,
				RepoError_3:      nil,
				RepoError_4:      nil,
				RepoError_5:      nil,
				RepoError_6:      nil,
				ExpectedReturn_1: false,
				ExpectedReturn_2: "",
				ExpectedError:    service.ErrCodeNotValid,
			},
		},
		{
			Name:               "Code Expired",
			VerifyRecoveryCode: tests.VerifyRecoveryCode_1,
			Repository: mock_struct.VerifyRecoveryCodeRepository{
				FindOneByEmail:         tests.User_1,
				RecoveryExists:         true,
				RecoveryFindOne:        tests.RecoveryCode_2,
				RecoveryTokenGenerator: "",
				RecoveryTokenInsertOne: nil,
			},
			ExpectedReturns: mock_struct.VerifyRecoveryCodeExpectedReturns{
				RepoError_1:      nil,
				RepoError_2:      nil,
				RepoError_3:      nil,
				RepoError_4:      nil,
				RepoError_5:      nil,
				RepoError_6:      nil,
				ExpectedReturn_1: false,
				ExpectedReturn_2: "",
				ExpectedError:    service.ErrCodeNotValid,
			},
		},
		{
			Name:               "FindOneByEmail RepositoryFailed",
			VerifyRecoveryCode: tests.VerifyRecoveryCode_1,
			Repository: mock_struct.VerifyRecoveryCodeRepository{
				FindOneByEmail:         nil,
				RecoveryExists:         false,
				RecoveryFindOne:        nil,
				RecoveryTokenGenerator: "",
				RecoveryTokenInsertOne: nil,
			},
			ExpectedReturns: mock_struct.VerifyRecoveryCodeExpectedReturns{
				RepoError_1:      utils.ErrRepositoryFailed,
				RepoError_2:      nil,
				RepoError_3:      nil,
				RepoError_4:      nil,
				RepoError_5:      nil,
				RepoError_6:      nil,
				ExpectedReturn_1: false,
				ExpectedReturn_2: "",
				ExpectedError:    utils.ErrRepositoryFailed,
			},
		},
		{
			Name:               "Exists RepositoryFailed",
			VerifyRecoveryCode: tests.VerifyRecoveryCode_1,
			Repository: mock_struct.VerifyRecoveryCodeRepository{
				FindOneByEmail:         tests.User_1,
				RecoveryExists:         false,
				RecoveryFindOne:        nil,
				RecoveryTokenGenerator: "",
				RecoveryTokenInsertOne: nil,
			},
			ExpectedReturns: mock_struct.VerifyRecoveryCodeExpectedReturns{
				RepoError_1:      nil,
				RepoError_2:      utils.ErrRepositoryFailed,
				RepoError_3:      nil,
				RepoError_4:      nil,
				RepoError_5:      nil,
				RepoError_6:      nil,
				ExpectedReturn_1: false,
				ExpectedReturn_2: "",
				ExpectedError:    utils.ErrRepositoryFailed,
			},
		},
		{
			Name:               "FindOne RepositoryFailed",
			VerifyRecoveryCode: tests.VerifyRecoveryCode_1,
			Repository: mock_struct.VerifyRecoveryCodeRepository{
				FindOneByEmail:         tests.User_1,
				RecoveryExists:         true,
				RecoveryFindOne:        nil,
				RecoveryTokenGenerator: "",
				RecoveryTokenInsertOne: nil,
			},
			ExpectedReturns: mock_struct.VerifyRecoveryCodeExpectedReturns{
				RepoError_1:      nil,
				RepoError_2:      nil,
				RepoError_3:      utils.ErrRepositoryFailed,
				RepoError_4:      nil,
				RepoError_5:      nil,
				RepoError_6:      nil,
				ExpectedReturn_1: false,
				ExpectedReturn_2: "",
				ExpectedError:    utils.ErrRepositoryFailed,
			},
		},
		{
			Name:               "UpdateOne RepositoryFailed",
			VerifyRecoveryCode: tests.VerifyRecoveryCode_1,
			Repository: mock_struct.VerifyRecoveryCodeRepository{
				FindOneByEmail:         tests.User_1,
				RecoveryExists:         true,
				RecoveryFindOne:        tests.RecoveryCode_1,
				RecoveryTokenGenerator: "",
				RecoveryTokenInsertOne: nil,
			},
			ExpectedReturns: mock_struct.VerifyRecoveryCodeExpectedReturns{
				RepoError_1:      nil,
				RepoError_2:      nil,
				RepoError_3:      nil,
				RepoError_4:      utils.ErrRepositoryFailed,
				RepoError_5:      nil,
				RepoError_6:      nil,
				ExpectedReturn_1: false,
				ExpectedReturn_2: "",
				ExpectedError:    utils.ErrRepositoryFailed,
			},
		},
		{
			Name:               "NewRecoveryCodeToken RepositoryFailed",
			VerifyRecoveryCode: tests.VerifyRecoveryCode_1,
			Repository: mock_struct.VerifyRecoveryCodeRepository{
				FindOneByEmail:         tests.User_1,
				RecoveryExists:         true,
				RecoveryFindOne:        tests.RecoveryCode_1,
				RecoveryTokenGenerator: "",
				RecoveryTokenInsertOne: nil,
			},
			ExpectedReturns: mock_struct.VerifyRecoveryCodeExpectedReturns{
				RepoError_1:      nil,
				RepoError_2:      nil,
				RepoError_3:      nil,
				RepoError_4:      nil,
				RepoError_5:      utils.ErrRepositoryFailed,
				RepoError_6:      nil,
				ExpectedReturn_1: false,
				ExpectedReturn_2: "",
				ExpectedError:    utils.ErrRepositoryFailed,
			},
		},
		{
			Name:               "InsertOne RepositoryFailed",
			VerifyRecoveryCode: tests.VerifyRecoveryCode_1,
			Repository: mock_struct.VerifyRecoveryCodeRepository{
				FindOneByEmail:         tests.User_1,
				RecoveryExists:         true,
				RecoveryFindOne:        tests.RecoveryCode_1,
				RecoveryTokenGenerator: tests.Token_1,
				RecoveryTokenInsertOne: nil,
			},
			ExpectedReturns: mock_struct.VerifyRecoveryCodeExpectedReturns{
				RepoError_1:      nil,
				RepoError_2:      nil,
				RepoError_3:      nil,
				RepoError_4:      nil,
				RepoError_5:      nil,
				RepoError_6:      utils.ErrRepositoryFailed,
				ExpectedReturn_1: false,
				ExpectedReturn_2: "",
				ExpectedError:    utils.ErrRepositoryFailed,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mockUserRepo.ExpectedCalls = nil
			mockRecoveryRepo.ExpectedCalls = nil
			mockRecoveryTokenRepo.ExpectedCalls = nil
			mockTokenGeneratorRepo.ExpectedCalls = nil

			mockUserRepo.On(tests.FIND_ONE_BY_EMAIL,
				mock.AnythingOfType(tests.STRING),
			).Return(tc.Repository.FindOneByEmail, tc.ExpectedReturns.RepoError_1)

			mockRecoveryRepo.On(tests.EXISTS,
				mock.AnythingOfType(tests.RECOVERY_CRITERIA_PTR),
			).Return(tc.Repository.RecoveryExists, tc.ExpectedReturns.RepoError_2)

			mockRecoveryRepo.On(tests.FIND_ONE,
				mock.AnythingOfType(tests.RECOVERY_CRITERIA_PTR),
			).Return(tc.Repository.RecoveryFindOne, tc.ExpectedReturns.RepoError_3)

			mockRecoveryRepo.On(tests.UPDATE_ONE,
				mock.AnythingOfType(tests.INT64),
				mock.AnythingOfType(tests.RECOVERY_DATA_UPDATE),
			).Return(tc.ExpectedReturns.RepoError_4)

			mockTokenGeneratorRepo.On(tests.NEW_RECOVERY_CODE_TOKEN,
				mock.AnythingOfType(tests.TIME),
				mock.AnythingOfType(tests.USER_MODEL),
			).Return(tc.Repository.RecoveryTokenGenerator, tc.ExpectedReturns.RepoError_5)

			mockRecoveryTokenRepo.On(tests.INSERT_ONE,
				mock.AnythingOfType(tests.RECOVERY_TOKEN_MODEL),
			).Return(tc.Repository.RecoveryTokenInsertOne, tc.ExpectedReturns.RepoError_6)

			result_1, result_2, err := RecoveryS.VerifyRecoveryCode(tc.VerifyRecoveryCode)

			mockUserRepo.AssertExpectations(t)
			mockRecoveryRepo.AssertExpectations(t)
			mockRecoveryTokenRepo.AssertExpectations(t)
			mockTokenGeneratorRepo.AssertExpectations(t)

			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)
			assert.Equal(t, tc.ExpectedReturns.ExpectedReturn_1, result_1)
			assert.Equal(t, tc.ExpectedReturns.ExpectedReturn_2, result_2)

		})
	}
}

func TestRecoveryCodeExpiry(t *testing.T) {
	testCases := []struct {
		Name string

		Repository struct {
			Find []model.Recovery
		}
		ExpectedReturns struct {
			RepoError_1   error
			RepoError_2   []error
			ExpectedError error
		}
	}{
		{
			Name: "Success",
			Repository: mock_struct.RecoveryCodeExpiryRepository{
				Find: []model.Recovery{
					*tests.RecoveryCode_1,
					*tests.RecoveryCode_2,
				},
			},
			ExpectedReturns: mock_struct.RecoveryCodeExpiryExpectedReturns{
				RepoError_1:   nil,
				RepoError_2:   []error{nil, nil},
				ExpectedError: nil,
			},
		},
		{
			Name: "Find RepositoryFailed",
			Repository: mock_struct.RecoveryCodeExpiryRepository{
				Find: []model.Recovery{},
			},
			ExpectedReturns: mock_struct.RecoveryCodeExpiryExpectedReturns{
				RepoError_1:   utils.ErrRepositoryFailed,
				RepoError_2:   []error{},
				ExpectedError: utils.ErrRepositoryFailed,
			},
		},
		{
			Name: "UpdateOne RepositoryFailed",
			Repository: mock_struct.RecoveryCodeExpiryRepository{
				Find: []model.Recovery{
					*tests.RecoveryCode_1,
					*tests.RecoveryCode_2,
				},
			},
			ExpectedReturns: mock_struct.RecoveryCodeExpiryExpectedReturns{
				RepoError_1:   nil,
				RepoError_2:   []error{utils.ErrRepositoryFailed, nil},
				ExpectedError: utils.ErrRepositoryFailed,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mockRecoveryRepo.ExpectedCalls = nil

			mockRecoveryRepo.On(tests.FIND,
				mock.AnythingOfType(tests.RECOVERY_CRITERIA_PTR),
			).Return(tc.Repository.Find, tc.ExpectedReturns.RepoError_1)

			for i := range tc.Repository.Find {
				mockRecoveryRepo.On(tests.UPDATE_ONE,
					mock.AnythingOfType(tests.INT64),
					mock.AnythingOfType(tests.RECOVERY_DATA_UPDATE),
				).Return(tc.ExpectedReturns.RepoError_2[i])
			}

			err := RecoveryS.RecoveryCodeExpiry()

			mockRecoveryRepo.AssertExpectations(t)

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

// 			Configuración del Mock para userRepository
// 			mockUserRepo.On(tests.EXISTS, mock.AnythingOfType(tests.USER_CRITERIA_PTR)).
// 				Return(tc.User.Exist, tc.ExpectedReturns.RepoError)

// 			Probar funcion
// 			err := s.UserExist(tc.UserID)
// 			mockUserRepo.AssertExpectations(t)

// 			Compara el error entregado, por el error esperado
// 			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)

// 		})
// 	}
// }
