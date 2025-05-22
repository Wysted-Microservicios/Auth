package mock_struct

import "github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"

type GetUserByIdRepository struct {
	FindOneByID *model.User
}

type GetUserByIdRepositoryExpectedReturns struct {
	RepoError      error
	ExpectedError  error
	ExpectedReturn *model.User
}
