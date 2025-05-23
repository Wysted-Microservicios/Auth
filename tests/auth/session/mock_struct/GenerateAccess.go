package mock_struct

import "github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"

type GenerateAccessParams struct {
	SessionToken string
	User         *model.User
}

type GenerateAccessRepository struct {
	SessionExists  int64
	TokenGenerator string
	AccessInsert   int64
}
type GenerateAccessExpectedReturns struct {
	RepoError_1    error
	RepoError_2    error
	RepoError_3    error
	ExpectedReturn string
	ExpectedError  error
}
