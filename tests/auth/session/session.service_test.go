package session

import (
	"os"
	"testing"

	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/dto"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/service"
	"github.com/CPU-commits/Template_Go-EventDriven/src/utils"
	"github.com/CPU-commits/Template_Go-EventDriven/tests"
	"github.com/CPU-commits/Template_Go-EventDriven/tests/auth/session/mock_struct"
	"github.com/CPU-commits/Template_Go-EventDriven/tests/mocks"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

var SessionS *service.SessionService

var mockSessionRepo *mocks.MockSessionRepository
var mockAccessRepo *mocks.MockAccessRepository
var mockTokenGenerator *mocks.MockTokenGenerator

func TestMain(m *testing.M) {

	mockSessionRepo = &mocks.MockSessionRepository{}
	mockAccessRepo = &mocks.MockAccessRepository{}
	mockTokenGenerator = &mocks.MockTokenGenerator{}

	SessionS = service.NewSessionService(mockSessionRepo, mockAccessRepo, mockTokenGenerator)

	code := m.Run()
	os.Exit(code)
}

func TestNewSession(t *testing.T) {
	testCases := []struct {
		Name   string
		Params struct {
			SessionDto dto.SessionDto
			AuthID     int64
			UserID     int64
		}
		Repository struct {
			NewSessionToken string
			InsertOne       int64
		}
		ExpectedReturns struct {
			RepoError_1    error
			RepoError_2    error
			ExpectedReturn string
			ExpectedError  error
		}
	}{
		{
			Name: "NewSession success",
			Params: mock_struct.NewSessionParams{
				SessionDto: tests.SessionDto,
				AuthID:     tests.ID_1,
				UserID:     tests.ID_1,
			},
			Repository: mock_struct.NewSessionRepository{
				NewSessionToken: tests.Token_1,
				InsertOne:       tests.Session_1.ID,
			},
			ExpectedReturns: mock_struct.NewSessionExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    nil,
				ExpectedReturn: tests.Token_1,
				ExpectedError:  nil,
			},
		},
		{
			Name: "NewSessionToken RepositoryFailed",
			Params: mock_struct.NewSessionParams{
				SessionDto: tests.SessionDto,
				AuthID:     tests.ID_1,
				UserID:     tests.ID_1,
			},
			Repository: mock_struct.NewSessionRepository{
				NewSessionToken: "",
				InsertOne:       0,
			},
			ExpectedReturns: mock_struct.NewSessionExpectedReturns{
				RepoError_1:    utils.ErrRepositoryFailed,
				RepoError_2:    nil,
				ExpectedReturn: "",
				ExpectedError:  utils.ErrRepositoryFailed,
			},
		},
		{
			Name: "InsertOne RepositoryFailed",
			Params: mock_struct.NewSessionParams{
				SessionDto: tests.SessionDto,
				AuthID:     tests.ID_1,
				UserID:     tests.ID_1,
			},
			Repository: mock_struct.NewSessionRepository{
				NewSessionToken: tests.Token_1,
				InsertOne:       0,
			},
			ExpectedReturns: mock_struct.NewSessionExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    utils.ErrRepositoryFailed,
				ExpectedReturn: "",
				ExpectedError:  utils.ErrRepositoryFailed,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mockTokenGenerator.ExpectedCalls = nil
			mockSessionRepo.ExpectedCalls = nil

			mockTokenGenerator.On(tests.NEW_SESSION_TOKEN,
				mock.AnythingOfType(tests.TIME),
				mock.AnythingOfType(tests.INT64),
			).Return(tc.Repository.NewSessionToken, tc.ExpectedReturns.RepoError_1)

			mockSessionRepo.On(tests.INSERT_ONE,
				mock.AnythingOfType(tests.SESSION_MODEL),
			).Return(tc.Repository.InsertOne, tc.ExpectedReturns.RepoError_2)

			result, err := SessionS.NewSession(tc.Params.SessionDto, tc.Params.AuthID, tc.Params.UserID)

			mockSessionRepo.AssertExpectations(t)
			mockTokenGenerator.AssertExpectations(t)

			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)
			assert.Equal(t, tc.ExpectedReturns.ExpectedReturn, result)
		})
	}
}

func TestRefreshSession(t *testing.T) {
	testCases := []struct {
		Name   string
		Params struct {
			SessionToken string
			UserID       int64
		}
		Repository struct {
			NewSessionToken string
		}
		ExpectedReturns struct {
			RepoError_1    error
			RepoError_2    error
			ExpectedReturn string
			ExpectedError  error
		}
	}{
		{
			Name: "RefreshSession success",
			Params: mock_struct.RefreshSessionParams{
				SessionToken: tests.Token_1,
				UserID:       tests.ID_1,
			},
			Repository: mock_struct.RefreshSessionRepository{
				NewSessionToken: tests.Token_1,
			},
			ExpectedReturns: mock_struct.RefreshSessionExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    nil,
				ExpectedReturn: tests.Token_1,
				ExpectedError:  nil,
			},
		},
		{
			Name: "NewSessionToken RepositoryFailed",
			Params: mock_struct.RefreshSessionParams{
				SessionToken: tests.Token_1,
				UserID:       tests.ID_1,
			},
			Repository: mock_struct.RefreshSessionRepository{
				NewSessionToken: "",
			},
			ExpectedReturns: mock_struct.RefreshSessionExpectedReturns{
				RepoError_1:    utils.ErrRepositoryFailed,
				RepoError_2:    nil,
				ExpectedReturn: "",
				ExpectedError:  utils.ErrRepositoryFailed,
			},
		},
		{
			Name: "Update RepositoryFailed",
			Params: mock_struct.RefreshSessionParams{
				SessionToken: tests.Token_1,
				UserID:       tests.ID_1,
			},
			Repository: mock_struct.RefreshSessionRepository{
				NewSessionToken: tests.Token_1,
			},
			ExpectedReturns: mock_struct.RefreshSessionExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    utils.ErrRepositoryFailed,
				ExpectedReturn: tests.Token_1,
				ExpectedError:  utils.ErrRepositoryFailed,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mockTokenGenerator.ExpectedCalls = nil
			mockSessionRepo.ExpectedCalls = nil

			mockTokenGenerator.On(tests.NEW_SESSION_TOKEN,
				mock.AnythingOfType(tests.TIME),
				mock.AnythingOfType(tests.INT64),
			).Return(tc.Repository.NewSessionToken, tc.ExpectedReturns.RepoError_1)

			mockSessionRepo.On(tests.UPDATE,
				mock.AnythingOfType(tests.SESSION_CRITERIA_PTR),
				mock.AnythingOfType(tests.SESSION_DATA_UPDATE),
			).Return(tc.ExpectedReturns.RepoError_2)

			result, err := SessionS.RefreshSession(tc.Params.SessionToken, tc.Params.UserID)

			mockTokenGenerator.AssertExpectations(t)
			mockSessionRepo.AssertExpectations(t)

			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)
			assert.Equal(t, tc.ExpectedReturns.ExpectedReturn, result)
		})
	}
}

func TestGenerateAccess(t *testing.T) {
	testCases := []struct {
		Name   string
		Params struct {
			SessionToken string
			User         *model.User
		}

		Repository struct {
			SessionExists  int64
			TokenGenerator string
			AccessInsert   int64
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
			Name: "GenerateAccess success",
			Params: mock_struct.GenerateAccessParams{
				SessionToken: tests.Token_1,
				User:         tests.User_1,
			},
			Repository: mock_struct.GenerateAccessRepository{
				SessionExists:  tests.ID_1,
				TokenGenerator: tests.Token_1,
				AccessInsert:   tests.ID_1,
			},
			ExpectedReturns: mock_struct.GenerateAccessExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    nil,
				RepoError_3:    nil,
				ExpectedReturn: tests.Token_1,
				ExpectedError:  nil,
			},
		},
		{
			Name: "Session not exists",
			Params: mock_struct.GenerateAccessParams{
				SessionToken: tests.Token_1,
				User:         tests.User_1,
			},
			Repository: mock_struct.GenerateAccessRepository{
				SessionExists:  0,
				TokenGenerator: "",
				AccessInsert:   0,
			},
			ExpectedReturns: mock_struct.GenerateAccessExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    nil,
				RepoError_3:    nil,
				ExpectedReturn: "",
				ExpectedError:  service.ErrSessionNotExists,
			},
		},
		{
			Name: "Exists RepositoryFailed",
			Params: mock_struct.GenerateAccessParams{
				SessionToken: tests.Token_1,
				User:         tests.User_1,
			},
			Repository: mock_struct.GenerateAccessRepository{
				SessionExists:  0,
				TokenGenerator: "",
				AccessInsert:   0,
			},
			ExpectedReturns: mock_struct.GenerateAccessExpectedReturns{
				RepoError_1:    utils.ErrRepositoryFailed,
				RepoError_2:    nil,
				RepoError_3:    nil,
				ExpectedReturn: "",
				ExpectedError:  utils.ErrRepositoryFailed,
			},
		},
		{
			Name: "NewAccessToken RepositoryFailed",
			Params: mock_struct.GenerateAccessParams{
				SessionToken: tests.Token_1,
				User:         tests.User_1,
			},
			Repository: mock_struct.GenerateAccessRepository{
				SessionExists:  tests.ID_1,
				TokenGenerator: "",
				AccessInsert:   0,
			},
			ExpectedReturns: mock_struct.GenerateAccessExpectedReturns{
				RepoError_1:    nil,
				RepoError_2:    utils.ErrRepositoryFailed,
				RepoError_3:    nil,
				ExpectedReturn: "",
				ExpectedError:  utils.ErrRepositoryFailed,
			},
		},
		{
			Name: "InsertOne RepositoryFailed",
			Params: mock_struct.GenerateAccessParams{
				SessionToken: tests.Token_1,
				User:         tests.User_1,
			},
			Repository: mock_struct.GenerateAccessRepository{
				SessionExists:  tests.ID_1,
				TokenGenerator: tests.Token_1,
				AccessInsert:   0,
			},
			ExpectedReturns: mock_struct.GenerateAccessExpectedReturns{
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
			mockSessionRepo.ExpectedCalls = nil
			mockTokenGenerator.ExpectedCalls = nil
			mockAccessRepo.ExpectedCalls = nil

			mockSessionRepo.On(tests.EXISTS,
				mock.AnythingOfType(tests.SESSION_CRITERIA_PTR),
			).Return(tc.Repository.SessionExists, tc.ExpectedReturns.RepoError_1)

			mockTokenGenerator.On(tests.NEW_ACCESS_TOKEN,
				mock.AnythingOfType(tests.TIME),
				mock.AnythingOfType(tests.USER_MODEL),
			).Return(tc.Repository.TokenGenerator, tc.ExpectedReturns.RepoError_2)

			mockAccessRepo.On(tests.INSERT_ONE,
				mock.AnythingOfType(tests.ACCESS_MODEL),
			).Return(tc.Repository.AccessInsert, tc.ExpectedReturns.RepoError_3)

			result, err := SessionS.GenerateAccess(tc.Params.SessionToken, tc.Params.User)

			mockSessionRepo.AssertExpectations(t)
			mockTokenGenerator.AssertExpectations(t)
			mockAccessRepo.AssertExpectations(t)

			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)
			assert.Equal(t, tc.ExpectedReturns.ExpectedReturn, result)
		})
	}
}
func TestCheckToken(t *testing.T) {
	testCases := []struct {
		Name  string
		Token string

		Exists int64

		ExpectedReturns struct {
			RepoError_1   error
			ExpectedError error
		}
	}{
		{
			Name:   "CheckToken success",
			Token:  tests.Token_1,
			Exists: 1,
			ExpectedReturns: mock_struct.CheckTokenExpectedReturns{
				RepoError_1:   nil,
				ExpectedError: nil,
			},
		},
		{
			Name:   "Token revoked or not exists",
			Token:  tests.Token_1,
			Exists: 0,
			ExpectedReturns: mock_struct.CheckTokenExpectedReturns{
				RepoError_1:   nil,
				ExpectedError: service.ErrTokenRevokedOrNotExists,
			},
		},
		{
			Name:   "Exists RepositoryFailed",
			Token:  tests.Token_1,
			Exists: 0,
			ExpectedReturns: mock_struct.CheckTokenExpectedReturns{
				RepoError_1:   utils.ErrRepositoryFailed,
				ExpectedError: utils.ErrRepositoryFailed,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mockAccessRepo.ExpectedCalls = nil

			mockAccessRepo.On(tests.EXISTS,
				mock.AnythingOfType(tests.ACCESS_CRITERIA_PTR),
			).Return(tc.Exists, tc.ExpectedReturns.RepoError_1)

			err := SessionS.CheckToken(tc.Token)

			mockAccessRepo.AssertExpectations(t)
			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)
		})
	}
}

func TestDeleteRevokedTokens(t *testing.T) {
	testCases := []struct {
		Name string

		ExpectedReturns struct {
			RepoError     error
			ExpectedError error
		}
	}{
		{
			Name: "DeleteRevokedTokens success",
			ExpectedReturns: mock_struct.DeleteRevokedTokensExpectedReturns{
				RepoError:     nil,
				ExpectedError: nil,
			},
		},
		{
			Name: "Delete RepositoryFailed",
			ExpectedReturns: mock_struct.DeleteRevokedTokensExpectedReturns{
				RepoError:     utils.ErrRepositoryFailed,
				ExpectedError: utils.ErrRepositoryFailed,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mockAccessRepo.ExpectedCalls = nil

			mockAccessRepo.On(tests.DELETE,
				mock.AnythingOfType(tests.ACCESS_CRITERIA_PTR),
			).Return(tc.ExpectedReturns.RepoError)

			err := SessionS.DeleteRevokedTokens()

			mockAccessRepo.AssertExpectations(t)

			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)

		})
	}
}

func TestDeleteExpiredSessions(t *testing.T) {
	testCases := []struct {
		Name string

		ExpectedReturns struct {
			RepoError     error
			ExpectedError error
		}
	}{
		{
			Name: "DeleteExpiredSessions success",
			ExpectedReturns: mock_struct.DeleteExpiredSessionsExpectedReturns{
				RepoError:     nil,
				ExpectedError: nil,
			},
		},
		{
			Name: "Delete RepositoryFailed",
			ExpectedReturns: mock_struct.DeleteExpiredSessionsExpectedReturns{
				RepoError:     utils.ErrRepositoryFailed,
				ExpectedError: utils.ErrRepositoryFailed,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			mockAccessRepo.ExpectedCalls = nil

			mockAccessRepo.On(tests.DELETE,
				mock.AnythingOfType(tests.ACCESS_CRITERIA_PTR),
			).Return(tc.ExpectedReturns.RepoError)

			err := SessionS.DeleteRevokedTokens()

			mockAccessRepo.AssertExpectations(t)

			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)
		})
	}
}
