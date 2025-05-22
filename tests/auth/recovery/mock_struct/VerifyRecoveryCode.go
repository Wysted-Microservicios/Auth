package mock_struct

import "github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"

type VerifyRecoveryCodeRepository struct {
	FindOneByEmail         *model.User
	RecoveryExists         bool
	RecoveryFindOne        *model.Recovery
	RecoveryTokenGenerator string
	RecoveryTokenInsertOne *model.RecoveryToken
}
type VerifyRecoveryCodeExpectedReturns struct {
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
