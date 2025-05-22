package mock_struct

import "github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"

type RecoveryCodeRepository struct {
	FindOneByEmail    *model.User
	RecoveryInsertOne *model.Recovery
}
type RecoveryCodeExpectedReturns struct {
	RepoError_1    error
	RepoError_2    error
	ExpectedReturn *model.Recovery
	ExpectedError  error
}
