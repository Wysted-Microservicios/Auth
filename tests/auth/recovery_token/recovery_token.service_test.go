package recovery_token

import (
	"os"
	"testing"

	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/service"
	"github.com/CPU-commits/Template_Go-EventDriven/src/utils"
	"github.com/CPU-commits/Template_Go-EventDriven/tests"
	"github.com/CPU-commits/Template_Go-EventDriven/tests/auth/recovery_token/mock_struct"
	"github.com/CPU-commits/Template_Go-EventDriven/tests/mocks"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

var RecoveryTokenS *service.RecoveryTokenService

var mockRecoveryToken *mocks.MockRecoveryTokenRepository
var mockTokenGenerator *mocks.MockTokenGenerator
var mockUserRepo *mocks.MockUserRepository

func TestMain(m *testing.M) {

	mockRecoveryToken = &mocks.MockRecoveryTokenRepository{}
	mockTokenGenerator = &mocks.MockTokenGenerator{}
	mockUserRepo = &mocks.MockUserRepository{}

	RecoveryTokenS = service.NewRecoveryTokenService(mockRecoveryToken, mockTokenGenerator, mockUserRepo)

	code := m.Run()
	os.Exit(code)
}

func TestNewRecoveryToken(t *testing.T) {
	testCases := []struct {
		Name string
		User model.User

		Repository struct {
			Exists            bool
			RecoveryCodeToken string
			InsertOne         *model.RecoveryToken
		}
		ExpectedReturns struct {
			RepoError_1    error
			RepoError_2    error
			RepoError_3    error
			ExpectedReturn string
			ExpectedError  error
		}
	}{
		{
			Name: "New Recovery Token success",
			User: *tests.User_1,
			Repository: mock_struct.NewRecoveryTokenRepository{
				Exists:            true,
				RecoveryCodeToken: tests.Token_1,
				InsertOne:         tests.RecoveryToken_1,
			},
			ExpectedReturns: mock_struct.NewRecoveryTokenExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    nil,
				RepoError_3:    nil,
				ExpectedReturn: tests.Token_1,
				ExpectedError:  nil,
			},
		},
		{
			Name: "User not found",
			User: *tests.User_1,
			Repository: mock_struct.NewRecoveryTokenRepository{
				Exists:            false,
				RecoveryCodeToken: "",
				InsertOne:         nil,
			},
			ExpectedReturns: mock_struct.NewRecoveryTokenExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    nil,
				RepoError_3:    nil,
				ExpectedReturn: "",
				ExpectedError:  service.ErrUserNotFound,
			},
		},
		{
			Name: "Exists RepositoryFailed",
			User: *tests.User_1,
			Repository: mock_struct.NewRecoveryTokenRepository{
				Exists:            false,
				RecoveryCodeToken: "",
				InsertOne:         nil,
			},
			ExpectedReturns: mock_struct.NewRecoveryTokenExpectedReturns{
				RepoError_1:    utils.ErrRepositoryFailed,
				RepoError_2:    nil,
				RepoError_3:    nil,
				ExpectedReturn: "",
				ExpectedError:  utils.ErrRepositoryFailed,
			},
		},
		{
			Name: "NewRecoveryCode RepositoryFailed",
			User: *tests.User_1,
			Repository: mock_struct.NewRecoveryTokenRepository{
				Exists:            true,
				RecoveryCodeToken: "",
				InsertOne:         nil,
			},
			ExpectedReturns: mock_struct.NewRecoveryTokenExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    utils.ErrRepositoryFailed,
				RepoError_3:    nil,
				ExpectedReturn: "",
				ExpectedError:  utils.ErrRepositoryFailed,
			},
		},
		{
			Name: "InsertOne RepositoryFailed",
			User: *tests.User_1,
			Repository: mock_struct.NewRecoveryTokenRepository{
				Exists:            true,
				RecoveryCodeToken: tests.Token_1,
				InsertOne:         nil,
			},
			ExpectedReturns: mock_struct.NewRecoveryTokenExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    nil,
				RepoError_3:    utils.ErrRepositoryFailed,
				ExpectedReturn: "",
				ExpectedError:  utils.ErrRepositoryFailed,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mockRecoveryToken.ExpectedCalls = nil
			mockTokenGenerator.ExpectedCalls = nil
			mockUserRepo.ExpectedCalls = nil

			mockUserRepo.On(tests.EXISTS,
				mock.AnythingOfType(tests.USER_CRITERIA_PTR),
			).Return(tc.Repository.Exists, tc.ExpectedReturns.RepoError_1)

			mockTokenGenerator.On(tests.NEW_RECOVERY_CODE_TOKEN,
				mock.AnythingOfType(tests.TIME),
				mock.AnythingOfType(tests.USER_MODEL),
			).Return(tc.Repository.RecoveryCodeToken, tc.ExpectedReturns.RepoError_2)

			mockRecoveryToken.On(tests.INSERT_ONE,
				mock.AnythingOfType(tests.RECOVERY_TOKEN_MODEL),
			).Return(tc.Repository.InsertOne, tc.ExpectedReturns.RepoError_3)

			result, err := RecoveryTokenS.NewRecoveryToken(tc.User)

			mockRecoveryToken.AssertExpectations(t)
			mockTokenGenerator.AssertExpectations(t)

			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)
			assert.Equal(t, tc.ExpectedReturns.ExpectedReturn, result)

		})
	}
}

func TestCheckToken(t *testing.T) {
	testCases := []struct {
		Name  string
		Token string

		Repository struct {
			Exists  bool
			FindOne *model.RecoveryToken
		}
		ExpectedReturns struct {
			RepoError_1    error
			RepoError_2    error
			ExpectedReturn model.RecoveryToken
			ExpectedError  error
		}
	}{
		{
			Name:  "CheckToken success",
			Token: tests.Token_1,
			Repository: mock_struct.CheckTokenRepository{
				Exists:  true,
				FindOne: tests.RecoveryToken_1,
			},
			ExpectedReturns: mock_struct.CheckTokenExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    nil,
				ExpectedReturn: *tests.RecoveryToken_1,
				ExpectedError:  nil,
			},
		},
		{
			Name:  "Token not valid",
			Token: tests.Token_1,
			Repository: mock_struct.CheckTokenRepository{
				Exists:  false,
				FindOne: nil,
			},
			ExpectedReturns: mock_struct.CheckTokenExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    nil,
				ExpectedReturn: model.RecoveryToken{},
				ExpectedError:  service.ErrTokenNotValid,
			},
		},
		{
			Name:  "Token not valid(Expired)",
			Token: tests.Token_1,
			Repository: mock_struct.CheckTokenRepository{
				Exists:  true,
				FindOne: tests.RecoveryToken_2,
			},
			ExpectedReturns: mock_struct.CheckTokenExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    nil,
				ExpectedReturn: model.RecoveryToken{},
				ExpectedError:  service.ErrTokenNotValid,
			},
		},
		{
			Name:  "Exists RepositoryFailed",
			Token: tests.Token_1,
			Repository: mock_struct.CheckTokenRepository{
				Exists:  false,
				FindOne: nil,
			},
			ExpectedReturns: mock_struct.CheckTokenExpectedReturns{
				RepoError_1:    utils.ErrRepositoryFailed,
				RepoError_2:    nil,
				ExpectedReturn: model.RecoveryToken{},
				ExpectedError:  utils.ErrRepositoryFailed,
			},
		},
		{
			Name:  "FindOne RepositoryFailed",
			Token: tests.Token_1,
			Repository: mock_struct.CheckTokenRepository{
				Exists:  true,
				FindOne: tests.RecoveryToken_1,
			},
			ExpectedReturns: mock_struct.CheckTokenExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    utils.ErrRepositoryFailed,
				ExpectedReturn: model.RecoveryToken{},
				ExpectedError:  utils.ErrRepositoryFailed,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			mockRecoveryToken.ExpectedCalls = nil

			mockRecoveryToken.On(tests.EXISTS,
				mock.AnythingOfType(tests.RECOVERY_TOKEN_CRITERIA_PTR),
			).Return(tc.Repository.Exists, tc.ExpectedReturns.RepoError_1)

			mockRecoveryToken.On(tests.FIND_ONE,
				mock.AnythingOfType(tests.RECOVERY_TOKEN_CRITERIA_PTR),
			).Return(tc.Repository.FindOne, tc.ExpectedReturns.RepoError_2)

			result, err := RecoveryTokenS.CheckToken(tc.Token)

			mockRecoveryToken.AssertExpectations(t)

			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)
			assert.Equal(t, tc.ExpectedReturns.ExpectedReturn, result)

		})
	}
}

func TestRecoveryTokenExpiry(t *testing.T) {
	testCases := []struct {
		Name       string
		Repository struct {
			Find []model.RecoveryToken
		}
		ExpectedReturns struct {
			RepoError_1   error
			RepoError_2   []error
			ExpectedError error
		}
	}{
		{
			Name: "RecoveryTokenExpiry success",
			Repository: mock_struct.RecoveryTokenExpiryRepository{
				Find: []model.RecoveryToken{
					*tests.RecoveryToken_1,
					*tests.RecoveryToken_2,
				},
			},
			ExpectedReturns: mock_struct.RecoveryTokenExpiryExpectedReturns{
				RepoError_1:   nil,
				RepoError_2:   []error{nil, nil},
				ExpectedError: nil,
			},
		},
		{
			Name: "Find RepositoryFailed",
			Repository: mock_struct.RecoveryTokenExpiryRepository{
				Find: []model.RecoveryToken{},
			},
			ExpectedReturns: mock_struct.RecoveryTokenExpiryExpectedReturns{
				RepoError_1:   utils.ErrRepositoryFailed,
				RepoError_2:   []error{},
				ExpectedError: utils.ErrRepositoryFailed,
			},
		},
		{
			Name: "UpdateOne RepositoryFailed",
			Repository: mock_struct.RecoveryTokenExpiryRepository{
				Find: []model.RecoveryToken{
					*tests.RecoveryToken_1,
					*tests.RecoveryToken_2,
				},
			},
			ExpectedReturns: mock_struct.RecoveryTokenExpiryExpectedReturns{
				RepoError_1:   nil,
				RepoError_2:   []error{utils.ErrRepositoryFailed, nil},
				ExpectedError: utils.ErrRepositoryFailed,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mockRecoveryToken.ExpectedCalls = nil

			mockRecoveryToken.On(tests.FIND,
				mock.AnythingOfType(tests.RECOVERY_TOKEN_CRITERIA_PTR),
			).Return(tc.Repository.Find, tc.ExpectedReturns.RepoError_1)

			for i := range tc.Repository.Find {
				mockRecoveryToken.On(tests.UPDATE_ONE,
					mock.AnythingOfType(tests.INT64),
					mock.AnythingOfType(tests.RECOVERY_TOKEN_DATA_UPDATE),
				).Return(tc.ExpectedReturns.RepoError_2[i])
			}

			err := RecoveryTokenS.RecoveryTokenExpiry()

			mockRecoveryToken.AssertExpectations(t)

			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)
		})
	}
}

// func TestRecoveryTokenExpiry(t *testing.T) {
// 	testCases := []struct {
// 		Name string
// 	}{}
// 	for _, tc := range testCases {
// 		t.Run(tc.Name, func(t *testing.T) {
//
// 		})
// 	}
// }
