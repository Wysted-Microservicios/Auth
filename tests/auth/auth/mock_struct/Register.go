package mock_struct

import "github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"

type RegisterRepository struct {
	Exists    bool
	InsertOne *model.User
}

type RegisterExpectedReturns struct {
	RepoError_1   error
	RepoError_2   error
	BusError      error
	ExpectedError error
}
