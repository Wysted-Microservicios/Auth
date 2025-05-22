package auth

import (
	"os"
	"testing"

	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/service"
	"github.com/CPU-commits/Template_Go-EventDriven/src/utils"
	"github.com/CPU-commits/Template_Go-EventDriven/tests"
	"github.com/CPU-commits/Template_Go-EventDriven/tests/auth/user/mock_struct"
	"github.com/CPU-commits/Template_Go-EventDriven/tests/mocks"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

var UserS *service.UserService

var mockUserRepo *mocks.MockUserRepository

func TestMain(m *testing.M) {
	mockUserRepo = &mocks.MockUserRepository{}

	UserS = service.NewUserService(mockUserRepo)

	code := m.Run()
	os.Exit(code)
}

func TestGetUserIDFromUsername(t *testing.T) {
	testCases := []struct {
		Name string
		// Parametros
		UserName string
		// Funciones
		User struct {
			FindOne *model.User
		}
		ExpectedReturns struct {
			RepoError      error
			ExpectedError  error
			ExpectedReturn int64
		}
	}{
		{
			Name:     "Username not exists",
			UserName: "WystedTest",
			User: mock_struct.GetUserIDFromUsernameRepository{
				FindOne: nil,
			},
			ExpectedReturns: mock_struct.GetUserIDFromUsernameExpectedReturns{
				RepoError:      nil,
				ExpectedError:  service.ErrUsernameNotExists,
				ExpectedReturn: 0,
			},
		},
		{
			Name:     "Username exists",
			UserName: "WystedTest",
			User: mock_struct.GetUserIDFromUsernameRepository{
				FindOne: tests.User_1,
			},
			ExpectedReturns: mock_struct.GetUserIDFromUsernameExpectedReturns{
				RepoError:      nil,
				ExpectedError:  nil,
				ExpectedReturn: 1,
			},
		},
		{
			Name:     "RepositoryFailed",
			UserName: "WystedTest",
			User: mock_struct.GetUserIDFromUsernameRepository{
				FindOne: nil,
			},
			ExpectedReturns: mock_struct.GetUserIDFromUsernameExpectedReturns{
				RepoError:      utils.ErrRepositoryFailed,
				ExpectedError:  utils.ErrRepositoryFailed,
				ExpectedReturn: 0,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Resetear expectativas de los mocks antes de cada test
			mockUserRepo.ExpectedCalls = nil

			mockUserRepo.On(tests.FIND_ONE,
				mock.AnythingOfType(tests.USER_CRITERIA_PTR),
				mock.AnythingOfType(tests.USER_FIND_ONE_OPTIOS_PTR),
			).Return(tc.User.FindOne, tc.ExpectedReturns.RepoError)

			result, err := UserS.GetUserIDFromUsername(tc.UserName)
			mockUserRepo.AssertExpectations(t)

			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)
			assert.Equal(t, tc.ExpectedReturns.ExpectedReturn, result)

		})
	}
}

func TestGetUserById(t *testing.T) {
	testCases := []struct {
		Name string
		// Parametros
		UserID int64
		User   struct {
			FindOneByID *model.User
		}
		ExpectedReturns struct {
			RepoError      error
			ExpectedError  error
			ExpectedReturn *model.User
		}
	}{
		{
			Name:   "User not found",
			UserID: 1,
			User: mock_struct.GetUserByIdRepository{
				FindOneByID: tests.User_1,
			},
			ExpectedReturns: mock_struct.GetUserByIdRepositoryExpectedReturns{
				RepoError:      nil,
				ExpectedError:  service.ErrUserNotFound,
				ExpectedReturn: nil,
			},
		},
		{
			Name:   "User found",
			UserID: 1,
			User: mock_struct.GetUserByIdRepository{
				FindOneByID: tests.User_1,
			},
			ExpectedReturns: mock_struct.GetUserByIdRepositoryExpectedReturns{
				RepoError:      nil,
				ExpectedError:  service.ErrUserNotFound,
				ExpectedReturn: nil,
			},
		},
		{
			Name:   "RepositoryFailed",
			UserID: 1,
			User: mock_struct.GetUserByIdRepository{
				FindOneByID: nil,
			},
			ExpectedReturns: mock_struct.GetUserByIdRepositoryExpectedReturns{
				RepoError:      utils.ErrRepositoryFailed,
				ExpectedError:  utils.ErrRepositoryFailed,
				ExpectedReturn: nil,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mockUserRepo.ExpectedCalls = nil

			mockUserRepo.On(tests.FIND_ONE_BY_ID,
				mock.AnythingOfType(tests.INT64),
			).Return(tc.ExpectedReturns.ExpectedReturn, tc.ExpectedReturns.RepoError)

			result, err := UserS.GetUserById(tc.UserID)
			mockUserRepo.AssertExpectations(t)

			assert.Equal(t, tc.ExpectedReturns.ExpectedError, err)
			assert.Equal(t, tc.ExpectedReturns.ExpectedReturn, result)

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
