package mock_struct

import "github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"

type LoginRepository struct {
	FindOneByUsername *model.Auth
	FindOneByEmail    *model.User
}
type LoginExpectedReturns struct {
	RepoError_1      error
	RepoError_2      error
	ExpectedReturn_1 *model.User
	ExpectedReturn_2 int64
	ExpectedError    error
}
