package mock_struct

import "github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"

type GetUserIDFromUsernameRepository struct {
	FindOne *model.User
}

type GetUserIDFromUsernameExpectedReturns struct {
	RepoError      error
	ExpectedError  error
	ExpectedReturn int64
}
